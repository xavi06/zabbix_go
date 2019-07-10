package zabbixgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// JSONRPCResponse struct
type JSONRPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Error   ZabbixError `json:"error"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

// JSONRPCRequest struct
type JSONRPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`

	// Zabbix 2.0:
	// The "user.login" method must be called without the "auth" parameter
	Auth string `json:"auth,omitempty"`
	ID   int    `json:"id"`
}

// ZabbixError struct
type ZabbixError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (z *ZabbixError) Error() string {
	return z.Data
}

// API struct
type API struct {
	url    string
	user   string
	passwd string
	id     int
	auth   string
	Client *http.Client
}

// NewAPI func
func NewAPI(server, user, passwd string) (*API, error) {
	return &API{server, user, passwd, 0, "", &http.Client{}}, nil
}

// GetAuth func
func (api *API) GetAuth() string {
	return api.auth
}

// ZabbixRequest func
func (api *API) ZabbixRequest(method string, data interface{}) (JSONRPCResponse, error) {
	// Setup our JSONRPC Request data
	id := api.id
	api.id = api.id + 1
	jsonobj := JSONRPCRequest{"2.0", method, data, api.auth, id}
	encoded, err := json.Marshal(jsonobj)

	if err != nil {
		return JSONRPCResponse{}, err
	}

	// Setup our HTTP request
	request, err := http.NewRequest("POST", api.url, bytes.NewBuffer(encoded))
	if err != nil {
		return JSONRPCResponse{}, err
	}
	request.Header.Add("Content-Type", "application/json-rpc")
	if api.auth != "" {
		// XXX Not required in practice, check spec
		//request.SetBasicAuth(api.user, api.passwd)
		//request.Header.Add("Authorization", api.auth)
	}

	// Execute the request
	response, err := api.Client.Do(request)
	if err != nil {
		return JSONRPCResponse{}, err
	}

	/**
	We can't rely on response.ContentLength because it will
	be set at -1 for large responses that are chunked. So
	we treat each API response as streamed data.
	*/
	var result JSONRPCResponse
	var buf bytes.Buffer

	_, err = io.Copy(&buf, response.Body)
	if err != nil {
		return JSONRPCResponse{}, err
	}

	json.Unmarshal(buf.Bytes(), &result)

	response.Body.Close()

	return result, nil
}

// Login func
func (api *API) Login() (bool, error) {
	params := make(map[string]string, 0)
	params["user"] = api.user
	params["password"] = api.passwd

	response, err := api.ZabbixRequest("user.login", params)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return false, err
	}

	if response.Error.Code != 0 {
		return false, &response.Error
	}
	if response.Result == nil {
		return false, errors.New("result is nil")
	}
	api.auth = response.Result.(string)
	return true, nil
}

// Logout func
func (api *API) Logout() (bool, error) {
	emptyparams := make(map[string]string, 0)
	response, err := api.ZabbixRequest("user.logout", emptyparams)
	if err != nil {
		return false, err
	}

	if response.Error.Code != 0 {
		return false, &response.Error
	}

	return true, nil
}

// Version func
func (api *API) Version() (string, error) {
	response, err := api.ZabbixRequest("APIInfo.version", make(map[string]string, 0))
	if err != nil {
		return "", err
	}

	if response.Error.Code != 0 {
		return "", &response.Error
	}

	return response.Result.(string), nil
}

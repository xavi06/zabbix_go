package zabbixgo

import (
	"encoding/json"
	"fmt"
)

// Alert type
type Alert map[string]interface{}

// Alert func
func (api *API) Alert(method string, data interface{}) {
	response, err := api.ZabbixRequest("alert."+method, data)
	if err != nil {
	}

	if response.Error.Code != 0 {
	}

	res, err := json.Marshal(response.Result)
	var ret []Alert
	err = json.Unmarshal(res, &ret)
	fmt.Println(ret, err)

}

package zabbixgo

import (
	"encoding/json"
	"fmt"
)

// Problem type
type Problem map[string]interface{}

// Problem func
func (api *API) Problem(method string, data interface{}) {
	response, err := api.ZabbixRequest("trigger."+method, data)
	if err != nil {
	}

	if response.Error.Code != 0 {
	}

	res, err := json.Marshal(response.Result)
	var ret []Problem
	err = json.Unmarshal(res, &ret)
	fmt.Println(ret, err)

}

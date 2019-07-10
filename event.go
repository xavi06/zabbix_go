package zabbixgo

import (
	"encoding/json"
	"fmt"
)

// Event type
type Event map[string]interface{}

// Event func
func (api *API) Event(method string, data interface{}) {
	response, err := api.ZabbixRequest("trigger."+method, data)
	if err != nil {
	}

	if response.Error.Code != 0 {
	}

	res, err := json.Marshal(response.Result)
	var ret []Event
	err = json.Unmarshal(res, &ret)
	fmt.Println(ret, err)

}

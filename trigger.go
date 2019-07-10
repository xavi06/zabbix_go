package zabbixgo

import (
	"encoding/json"
)

type TriggerHost struct {
	Hostid string
	Host   string
	Name   string
}

type Trigger struct {
	Description string
	Lastchange  string
	Priority    string
	Triggerid   string
	Url         string
	Hosts       []TriggerHost
}

// Trigger type
//type Trigger map[string]interface{}

// Trigger func
func (api *API) Trigger(method string, data interface{}) ([]Trigger, error) {
	response, err := api.ZabbixRequest("trigger."+method, data)
	if err != nil {
		return nil, err
	}

	if response.Error.Code != 0 {
		return nil, &response.Error
	}

	res, err := json.Marshal(response.Result)
	var ret []Trigger
	err = json.Unmarshal(res, &ret)
	return ret, err

}

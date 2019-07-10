package zabbixgo

import "encoding/json"

type ZabbixHost map[string]interface{}

// Host func
func (api *API) Host(method string, data interface{}) ([]ZabbixHost, error) {
	response, err := api.ZabbixRequest("host."+method, data)
	if err != nil {
		return nil, err
	}

	if response.Error.Code != 0 {
		return nil, &response.Error
	}

	res, err := json.Marshal(response.Result)
	var ret []ZabbixHost
	err = json.Unmarshal(res, &ret)
	return ret, nil
}

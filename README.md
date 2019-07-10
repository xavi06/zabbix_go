# zabbix_go
zabbix golang api

#### Usage:
```bash
 go get github.com/xavi06/zabbix_go
```
```go
package main

import (
    "fmt"

    "github.com/xavi06/zabbix_go"
)

func main() {
    api, err := zabbixgo.NewAPI("http://wg.zabbix.fccs.cn/api_jsonrpc.php", "admin", "admin123")
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = api.Login()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Connected to API")
    params := make(map[string]interface{}, 0)
    params["output"] = []string{"triggerid", "description", "priority", "lastchange", "url"}
    //params["output"] = "extend"
    params["selectHosts"] = []string{"name", "host"}
    params["sortfield"] = "lastchange"
    params["sortorder"] = "DESC"
    params["limit"] = "1"
    filter := make(map[string]string, 0)
    filter["value"] = "1"
    params["filter"] = filter
    res, _ := api.Trigger("get", params)
    fmt.Println(res)
}
```

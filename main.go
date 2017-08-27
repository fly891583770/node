package main

import (
	"fmt"
	"time"

	"node/utils"
	//"node/plugins"
)

func main() {
	fmt.Println("hello world")
	//	settings := common.GetSettings()
	//	fmt.Println(settings.Get("INFLUXDB_HOST"))
	//	pm := plugins.PluginManager{}
	//	plugin := plugins.EventHandler{}
	//	plugins := make([]common.BasePlugin, 0)
	//	plugins = append(plugins, &plugin)
	//	pm.RegPlugins(plugins)
	//	fmt.Println(pm)
	client, err := utils.GetInfluxDBWriteClient()
	fmt.Println(client, err)
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	ok, err := utils.WriteData(client, "cpuusage", tags, fields, time.Now())
	fmt.Println(ok, err)
}

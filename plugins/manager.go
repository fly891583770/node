package plugins

import (
	"fmt"
	"log"
	"sync"

	"node/common"
)

type PluginManager struct {
	mutex   sync.Mutex
	plugins map[string]common.BasePlugin
}

func (pm *PluginManager) RegPlugins(plugins []common.BasePlugin) []error {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	ret := make([]error, 0, len(plugins))
	if pm.plugins == nil {
		pm.plugins = map[string]common.BasePlugin{}
	}

	for _, plugin := range plugins {
		name := plugin.GetPluginName()

		if _, found := pm.plugins[name]; found {
			ret = append(ret, fmt.Errorf("plugin %s: registered more than once", name))
			continue
		}
		err := plugin.Init()
		if err != nil {
			ret = append(ret, fmt.Errorf("plugin %s: init failed, %s", name, err.Error()))
			continue
		}
		pm.plugins[name] = plugin
		log.Printf("success loaded volume plugin %s", name)
	}
	return ret
}

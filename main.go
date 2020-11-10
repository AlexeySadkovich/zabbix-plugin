package main

import (
	"zabbix.com/pkg/plugin"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx Plugin.ContextProvider) (res interfase{}, err error) {
	// TODO
	return 
}

func init() {
	Plugin.RegisterMetrics(&impl, "PluginName", "key", "My zabbix plugin")
}
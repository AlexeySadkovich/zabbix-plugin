package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"zabbix.com/pkg/plugin"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx Plugin.ContextProvider) (res interface{}, err error) {
	if len(params) != 1 {
		return nil, errors.New("Wrong parameters")
	}

	res, err := p.httpClient.Get(fmt.Sprintf("https://wttr.in/~%s&format=%%t", params[0]))
	if err != nil {
		return nil, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return string(temp)[0 : len(temp)-4], nil
}

func init() {
	Plugin.RegisterMetrics(&impl, "Weather", "weather.temp", "Returns Celsius temperature.")
}

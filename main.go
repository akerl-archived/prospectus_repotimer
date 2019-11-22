package main

import (
	"github.com/akerl/prospectus/plugin"
)

type config struct {
	Days int `json:days`
}

type timerPlugin struct {
	Config config
}

func (p timerPlugin) GetConfigPointer() interface{} {
	return &p.Config
}

func (p timerPlugin) Load(input plugin.LoadInput) plugin.AttributeSet {
	return plugin.AttributeSet{plugin.Attribute{Name: "repotimer"}}
}

func (p timerPlugin) Check(input plugin.Attribute) plugin.Result {
	return plugin.Result{
		Actual:   "3",
		Expected: "1, 2, 3, 4, 5, 6, 7",
		Matches:  true,
	}
}

func (p timerPlugin) Fix(input plugin.Result) plugin.Result {
	return input
}

func main() {
	p := timerPlugin{}
	err := plugin.Start(p)
	if err != nil {
		panic(err)
	}
}

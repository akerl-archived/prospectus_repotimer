package main

import (
	"fmt"

	"github.com/akerl/prospectus/plugin"
)

type config struct {
	Days int `json:days`
}

func main() {
	var c config
	err := plugin.LoadPluginConfig(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", c)
}

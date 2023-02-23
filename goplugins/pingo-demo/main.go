package main

import (
	"fmt"

	"github.com/dullgiulio/pingo"
)

func main() {
	pgo := pingo.NewPlugin("tcp", "plugins/plugins")
	pgo.Start()

	defer pgo.Stop()

	var res string

	if err := pgo.Call("Plugin.Runing", "https://www.baidu.com", &res); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("demo")
	}
}

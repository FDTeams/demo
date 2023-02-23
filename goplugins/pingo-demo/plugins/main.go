package main

import (
	"fmt"

	"github.com/dullgiulio/pingo"
)

type Plugin struct{}

func (p *Plugin) Runing(target string, msg *string) error {
	*msg = "hello world " + target
	fmt.Println("the target from ", target)
	return nil
}

func main() {
	plugin := &Plugin{}
	pingo.Register(plugin)
	pingo.Run()
}

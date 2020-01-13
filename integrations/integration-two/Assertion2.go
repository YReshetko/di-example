package integration_two

import (
	"fmt"
	"github.com/YReshetko/di-example/registry"
)

func init() {
	registry.AddIntegrations("int2", &a2{})
}

type a2 struct {}

func (a *a2) Process() {
	fmt.Println("Integration two, assertion2")
}

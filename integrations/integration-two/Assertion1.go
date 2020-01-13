package integration_two

import (
	"fmt"
	"github.com/YReshetko/di-example/registry"
)

func init() {
	registry.AddIntegrations("int2", &a1{})
}

type a1 struct {}

func (a *a1) Process() {
	fmt.Println("Integration two, assertion1")
}

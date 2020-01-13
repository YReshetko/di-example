package integrtion_one

import (
	"fmt"
	"github.com/YReshetko/di-example/registry"
)

func init() {
	registry.AddIntegrations("int1", &a1{}, &a2{})
}

type a1 struct {}

func (a *a1) Process() {
	fmt.Println("Integration one, assertion1")
}

type a2 struct {}

func (a *a2) Process() {
	fmt.Println("Integration one, assertion2")
}

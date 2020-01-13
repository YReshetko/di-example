package main

import (
	"fmt"
	"github.com/YReshetko/di-example/registry"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("No integrtion is selected")
	}

	key := args[0]
	p, err := registry.GetIntegrations(key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Process some common assertions")

	for _, v := range p {
		v.Process()
	}
}

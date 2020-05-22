package main

import (
	"context"
	"fmt"
	"plugin"

	"github.com/aws/aws-lambda-go/lambda"
)

func CallLayer(p *plugin.Plugin, funcName string) string {
	_layer, err := p.Lookup(funcName)
	if err != nil {
		panic(err)
	}
	return _layer.(func() string)()
}

type MyEvent struct {
	Func string `json:"func"`
}

func Handler(ctx context.Context, event MyEvent) (string, error) {
	fmt.Printf("Execute func is : %v", event.Func)

	pluginModule, err := plugin.Open("/opt/layer.so")
	if err != nil {
		panic(err)
	}
	return CallLayer(pluginModule, event.Func), nil
}

func main() {
	lambda.Start(Handler)
}

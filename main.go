package main

import (
	"fmt"

	"github.com/tomassirio/goModules/controller"
	"github.com/tomassirio/goModules/model"
	"github.com/tomassirio/goModules/view"
)

func main() {
	m := &model.ModelStruct{Feature: "A Feature"}
	fmt.Printf("Our Model Feature looks like this '%s'\n", m.Feature)
	view.ViewFunction()
	controller.ControllerFunction()
}

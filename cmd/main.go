package main

import (
	"fmt"
	"github.com/qorganbek/kwaka_test_assignment/internal/app"
	"github.com/qorganbek/kwaka_test_assignment/internal/config"
)

func main() {

	cfg, err := config.InitConfig("config.yaml")

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)

	if err != nil {
		panic(err)
	}
}

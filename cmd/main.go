package cmd

import (
	"fmt"
	"kwaka_test/internal/app"
	"kwaka_test/internal/config"
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

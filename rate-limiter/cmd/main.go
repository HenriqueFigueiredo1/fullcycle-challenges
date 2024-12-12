package main

import (
	"fmt"
	"net/http"

	"github.com/HenriqueFigueiredo1/goexpert-rate-limiter/cmd/config"
	"github.com/HenriqueFigueiredo1/goexpert-rate-limiter/cmd/router"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		panic(fmt.Sprintf("error starting configs: %s", err.Error()))
	}

	r := router.StartTestRoutes(cfg)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic("error on listen and serve: " + err.Error())
	}
}

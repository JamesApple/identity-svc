package main

import (
	"root/cfg"
	"root/web"
)

func main() {

	app, err := cfg.NewApp(cfg.MakeConfig())
	if err != nil {
		panic(err)
	}

	web.Server(app)
}

package main

import (
	"flag"
	"fmt"

	"goshorten/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	fmt.Println(cfg.Server)

	//checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

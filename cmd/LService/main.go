package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/F7icK/LService/internal/app/apiserver"
)

var (
	configPath string
)

func init()  {
	flag.StringVar(&configPath, "config-path", "configs/lservice.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
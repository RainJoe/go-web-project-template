package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/RainJoe/go-template/pkg/adding"
	"github.com/RainJoe/go-template/pkg/config"
	"github.com/RainJoe/go-template/pkg/http/rest"
	"github.com/RainJoe/go-template/pkg/listing"
	"github.com/RainJoe/go-template/pkg/storage/postgres"
)

var tomlfile string

func init() {
	flag.StringVar(&tomlfile, "toml", "../../pkg/config/dev.toml", "TOML config file")
}

func main() {
	flag.Parse()
	//parse toml config file
	var conf config.TOMLConfig
	if _, err := toml.DecodeFile(tomlfile, &conf); err != nil {
		panic(err)
	}

	//init db
	s, err := postgres.NewStorage(conf)
	if err != nil {
		panic(err)
	}

	//init services
	adder := adding.NewService(s)
	lister := listing.NewService(s)

	//init router
	router := rest.Handler(adder, lister)

	//start server
	router.Run(conf.Port)
}

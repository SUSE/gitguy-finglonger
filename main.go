package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/SUSE/gitguy-finglonger/pkg/config"
	"github.com/SUSE/gitguy-finglonger/pkg/github/api"
)

var (
	configFile string
)

func main() {
	flag.StringVar(&configFile, "config", "config.devel.yml", "Set the path to the file with the configuration for the project.")
	flag.Parse()

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	conf := new(config.Config)
	err = yaml.Unmarshal(b, conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	a := api.NewAPI(conf)
	log.Fatal(a.ListenAndServe())
}

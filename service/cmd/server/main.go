package main

import (
	"fmt"
	"github.com/linger1216/go-front/service/svc/server"
	"github.com/linger1216/go-front/utils/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	configFilename = kingpin.Flag("conf", "yaml config file name").Short('c').
		Default("conf/config.yaml").String()
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()

	if _, err := os.Stat(*configFilename); err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("%s not found", *configFilename))
		} else {
			panic(err)
		}
	}

	reader := config.NewYamlReader(*configFilename)
	server.Run(reader)
}

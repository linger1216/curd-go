package main

import (
	"github.com/linger1216/go-front/echo-service/svc/server"
	"github.com/linger1216/go-front/utils/config"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configFilename = kingpin.Flag("conf", "yaml config file name").Short('c').
		Default("conf/config.yaml").String()
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()
	reader := config.NewYamlReader(*configFilename)
	server.Run(reader)
}

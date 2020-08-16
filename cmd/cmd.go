package cmd

import (
	"flag"
	"log"

	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/provider/process"
	"github.com/q8s-io/mystique/pkg/provider/spider"
)

var confPath = flag.String("conf", "./configs/pro.yaml", "The path of config.")

// 服务类型
var serverType = flag.String("type", "sidecar", "The type of server.")

func Run() {
	flag.Parse()
	// init
	process.Init(*confPath)
	// app
	RunApp(*serverType)
}

func RunApp(serverType string) {
	switch serverType {
	case "sidecar":
		RunSidecar()
	default:
		log.Println(serverType)
	}
}

func RunSidecar() {
	switch model.Config.App.Type {
	case "spider":
		go spider.Run()
	default:
		log.Println(model.Config.App.Type)
	}

	spider.Signal()
}

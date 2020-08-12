package cmd

import (
	"flag"
	"log"

	"github.com/q8s-io/mystique/pkg/provider/process"
	"github.com/q8s-io/mystique/pkg/provider/sidecar"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The path of config.")

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
	go sidecar.Run()
	sidecar.Signal()
}

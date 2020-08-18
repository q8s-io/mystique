package spider

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
)

func Signal() {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Println("signal cancelled")

		if model.Config.Input.Enable {
			// 处理业务逻辑
			close(kafka.Queue)
			// wait for kafka comsumer commit
			time.Sleep(1 * time.Second)

			close(StdinQueue)
		}

		if model.Config.Output.Enable {
			close(StdoutNativeDataQueue)
			close(StdoutQueue)
		}

		log.Println("server stop")
		os.Exit(0)
	}
}

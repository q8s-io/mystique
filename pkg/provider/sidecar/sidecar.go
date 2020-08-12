package sidecar

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
	"github.com/q8s-io/mystique/pkg/repository"
)

var StdinQueue chan []byte
var StdoutQueue chan []byte
var StderrQueue chan []byte

func runInLinux(cmd string) (string, error) {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(result)), err
	}
}

func GetPid(serverName string) string {
	a := `ps -x | grep "` + serverName + `" | head -1 | awk '{print $1}'`
	pid, err := runInLinux(a)
	if err != nil {
		log.Println(err)
		return ""
	} else {
		return pid
	}
}

func Run() {
	time.Sleep(5 * time.Second)

	InputConfig := model.Config.Input
	OutputConfig := model.Config.Output

	StdinQueue = make(chan []byte, 0)
	StdoutQueue = make(chan []byte, 0)

	pid := GetPid(model.Config.App.ProcessName)
	log.Println(pid)

	kafka.InitConsumer(InputConfig.BrokerList, "mystique")
	repository.ConsumerMsg(InputConfig.Topic)

	kafka.InitSyncProducer(OutputConfig.BrokerList)
	go func() {
		for msg := range StdoutQueue {
			repository.ProducerMsg(OutputConfig.Topic, msg)
		}
	}()

	go SinkToQueueFromStdout(pid)

	go SinkToStdinFromQueue(pid)

	for msg := range kafka.Queue {
		StdinQueue <- msg
	}
}

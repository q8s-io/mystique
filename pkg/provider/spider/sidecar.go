package spider

import (
	"log"
	"os/exec"
	"strings"

	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
	"github.com/q8s-io/mystique/pkg/repository"
)

var StdinQueue chan []byte
var StdoutQueue chan []byte
var StderrQueue chan []byte

func Run() {
	InputConfig := model.Config.Input
	OutputConfig := model.Config.Output

	StdinQueue = make(chan []byte, 0)
	StdoutQueue = make(chan []byte, 0)
	StderrQueue = make(chan []byte, 0)

	if InputConfig.Enable && OutputConfig.Enable {
		InputAndOutput(InputConfig, OutputConfig)
	} else if InputConfig.Enable && !OutputConfig.Enable {
		InputOnly(InputConfig)
	} else if !InputConfig.Enable && OutputConfig.Enable {
		OutputOnly(OutputConfig)
	} else {
		SelfOnly()
	}
}

func InputAndOutput(inputConfig model.Input, outputConfig model.Output) {
	pid := GetPid(model.Config.App.ProcessorsName)
	log.Println(pid)

	go SinkToQueueFromStdout(pid)

	kafka.InitSyncProducer(outputConfig.Kafka.Broker)
	go func() {
		for msg := range StdoutQueue {
			repository.ProducerMsg(inputConfig.Kafka.Topic, msg)
		}
	}()

	go SinkToStdinFromQueue(pid)

	kafka.InitConsumer(inputConfig.Kafka.Broker, inputConfig.Kafka.ConsumerGroup)
	repository.ConsumerMsg(inputConfig.Kafka.Topic)

	for msg := range kafka.Queue {
		StdinQueue <- msg
	}
}

func InputOnly(inputConfig model.Input) {

}

func OutputOnly(outputConfig model.Output) {

}

func SelfOnly() {

}

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

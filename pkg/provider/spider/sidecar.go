package spider

import (
	"log"

	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
	"github.com/q8s-io/mystique/pkg/repository"
)

var StdinQueue chan []byte
var StdoutQueue chan model.StdoutData

func Run() {
	InputConfig := model.Config.Input
	OutputConfig := model.Config.Output

	StdinQueue = make(chan []byte, 0)
	StdoutQueue = make(chan model.StdoutData, 1000)

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
	runStdout(pid, outputConfig)
	runStdin(pid, inputConfig)
}

func InputOnly(inputConfig model.Input) {
	pid := GetPid(model.Config.App.ProcessorsName)
	runStdin(pid, inputConfig)
}

func OutputOnly(outputConfig model.Output) {
	pid := GetPid(model.Config.App.ProcessorsName)
	runStdout(pid, outputConfig)
}

func SelfOnly() {

}

func runStdin(pid string, inputConfig model.Input) {
	kafka.InitConsumer(inputConfig.Kafka.Broker, inputConfig.Kafka.ConsumerGroup)

	go SinkToStdinFromQueue(pid)

	go repository.ConsumerMsg(inputConfig.Kafka.Topic)

	for msg := range kafka.Queue {
		StdinQueue <- msg
	}
}

func runStdout(pid string, outputConfig model.Output) {
	kafka.InitSyncProducer(outputConfig.Kafka.Broker)

	go SinkToQueueFromStdout(pid)
	go SinkToQueueFromStderr(pid)

	go func() {
		for msg := range StdoutQueue {
			switch msg.Type {
			case model.ConfigTypeKafka:
				repository.ProducerMsg(msg.Config, msg.Data)
			case model.ConfigTypeQbus:
				repository.ProducerMsg(msg.Config, msg.Data)
			default:
				log.Println(msg)
			}
		}
	}()
}

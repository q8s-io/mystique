package repository

import (
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
)

func ProducerMsg(topic string, msg []byte) {
	kafka.SyncProducerSendMsg(topic, string(msg))
}

func ConsumerMsg(topic string) {
	go kafka.ConsumerMsg(topic)
}

package repository

import (
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
)

func ProducerMsg(topic string, msg string) {
	kafka.SyncProducerSendMsg(topic, msg)
}

func ConsumerMsg(topic string) {
	kafka.ConsumerMsg(topic)
}

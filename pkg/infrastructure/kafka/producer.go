package kafka

import (
	"github.com/Shopify/sarama"

	"github.com/q8s-io/mystique/pkg/infrastructure/xray"
)

var SyncProducer sarama.SyncProducer
var syncProducerErr error

func InitSyncProducer(brokerList []string) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 1
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	SyncProducer, syncProducerErr = sarama.NewSyncProducer(brokerList, config)
	if syncProducerErr != nil {
		xray.ErrMini(syncProducerErr)
	}
}

func SyncProducerSendMsg(topic string, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := SyncProducer.SendMessage(msg)

	if err != nil {
		xray.ErrMini(err)
	}
}

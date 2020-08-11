package repository

import (
	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
)

// func SendMsgJobImageAnalyzer(jobScannerMsg *model.JobScannerMsg) {
// 	kafka.SyncProducerSendMsg("analyzer", jobScannerMsg)
// }

func ConsumerMsg() {
	go kafka.ConsumerMsg(model.Config.Kafka.Topic)
}

package sidecar

import (
	"log"

	"github.com/q8s-io/mystique/pkg/infrastructure/kafka"
	"github.com/q8s-io/mystique/pkg/repository"
)

var StdinQueue chan []byte
var StdoutQueue chan []byte
var StderrQueue chan []byte

func Run() {
	StdinQueue = make(chan []byte, 0)

	// consumer msg from mq
	repository.ConsumerMsg()

	go SinkToStdinFromQueue()

	for msg := range kafka.Queue {
		log.Printf("consumer msg from kafka %s", msg)
		StdinQueue <- msg
	}
}

// jobScannerMsg := new(model.JobScannerMsg)
// _ = json.Unmarshal(msg, &jobScannerMsg)

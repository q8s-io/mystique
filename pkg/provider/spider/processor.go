package spider

import (
	"fmt"
	"log"
	"strings"

	"github.com/q8s-io/mystique/pkg/entity/model"
)

func processOutputLine(line string) {
	outputData := strings.Split(line, "\t")
	var stdoutData model.StdoutData
	// data
	if outputData[0] == model.OutputTypeKafka {
		outputType := outputData[1]
		outputConfig := outputData[2]
		tempStr := fmt.Sprintf("output\t%s\t%s\t", outputType, outputConfig)
		realStr := string([]rune(line)[len([]rune(line))-(len(line)-len(tempStr)):])
		stdoutData.Type = outputType
		stdoutData.Config = outputConfig
		stdoutData.Data = realStr
		// log
	} else {
		stdoutData.Type = model.ConfigTypeKafka
		stdoutData.Config = model.Config.Output.Kafka.Topic
		stdoutData.Data = line
	}
	StdoutQueue <- stdoutData
	log.Println("read data from fd succeed.")
}

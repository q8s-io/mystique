package model

const (
	OutputTypeKafka string = "output"
	ConfigTypeKafka string = "kafka"
	ConfigTypeQbus  string = "qbus"
)

type StdoutData struct {
	Type   string
	Config string
	Data   string
}

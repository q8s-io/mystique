package model

const (
	OutputTypeData   string = "output"
	ConfigTypeKafka  string = "kafka"
	ConfigTypeQbus   string = "qbus"
	ConfigTypeScribe string = "scribe"
)

type StdoutNativeData struct {
	Data string
}

type StdoutData struct {
	Type   string
	Config string
	Data   string
}

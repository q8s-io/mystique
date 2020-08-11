package model

var Config Runtime

type Runtime struct {
	Kafka kafka
}

type kafka struct {
	BrokerList []string `toml:"broker"`
	Topic      string   `toml:"topic"`
}

package model

var Config Runtime

type Runtime struct {
	App    App
	Input  Input
	Output Output
}

type App struct {
	Type           string `yaml:"type"`
	ProcessorsName string `yaml:"processors_name"`
}

type Input struct {
	Enable bool  `yaml:"enable"`
	Kafka  Kafka `yaml:"kafka"`
}

type Output struct {
	Enable bool  `yaml:"enable"`
	Kafka  Kafka `yaml:"kafka"`
}

type Kafka struct {
	Broker        []string `yaml:"broker"`
	Topic         string   `yaml:"topic"`
	ConsumerGroup string   `yaml:"consumer_group"`
}

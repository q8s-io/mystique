package model

var Config Runtime

type Runtime struct {
	App    app
	Input  input
	Output output
}

type app struct {
	ProcessName string `toml:"process"`
}

type input struct {
	BrokerList []string `toml:"broker"`
	Topic      string   `toml:"topic"`
}

type output struct {
	BrokerList []string `toml:"broker"`
	Topic      string   `toml:"topic"`
}

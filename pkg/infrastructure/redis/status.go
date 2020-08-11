package redis

import (
	"github.com/q8s-io/mystique/pkg/infrastructure/xray"
)

func Status() {
	pong, err := Client.Ping().Result()
	if err != nil || pong != "PONG" {
		//xray.ErrMini(err)
		panic(xray.ErrMiniInfo(err))
	}
}

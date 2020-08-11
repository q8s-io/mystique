package mysql

import (
	"github.com/q8s-io/mystique/pkg/infrastructure/xray"
)

func Status() {
	pingErr := Client.DB().Ping()
	if pingErr != nil {
		panic(xray.ErrMiniInfo(pingErr))
	}
}

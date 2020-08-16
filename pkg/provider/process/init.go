package process

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/q8s-io/mystique/pkg/entity/model"
	"github.com/q8s-io/mystique/pkg/infrastructure/qlog"
)

func Init(confPath string) {
	// init config
	config, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = yaml.Unmarshal(config, &model.Config)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(model.Config)

	// init log
	qlog.InitLog()
}

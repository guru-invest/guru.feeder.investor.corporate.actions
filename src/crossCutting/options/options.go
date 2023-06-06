package options

import (
	"encoding/json"

	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
	log "github.com/sirupsen/logrus"
)

var OPTIONS Options

type Options struct {
	DATABASE  DatabaseOption  `json:"database"`
	ENDPOINTS EndPointsOption `json:"endpoints"`
}

func (c *Options) Load() {
	config := c.getConfig()
	kv := KeyValuePairOption{}
	if err := json.Unmarshal(config, &kv); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while in parse config json")
	}

	if err := json.Unmarshal(kv.DecodeString(), &c); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while decoding config")
	}

}

func (c *Options) getConfig() []byte {
	client := http_connector.HttpClient{}
	consulURL := ConsulOption{}.Get()
	res, err := client.Get(consulURL)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while get configurations")
	}
	return res
}

package options

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type ConsulOption struct {
	URL     string
	Port    string
	Context string
	Service string
}

func (t ConsulOption) Get() string {
	consulOption := ConsulOption{
		URL:     os.Getenv("CONSUL_URL"),
		Port:    os.Getenv("CONSUL_PORT"),
		Context: os.Getenv("CONSUL_CONTEXT"),
		Service: os.Getenv("CONSUL_SERVICE"),
	}

	if consulOption.URL == "" {
		log.Fatal("Error retrieving consul URL from env variables")
	}

	if consulOption.Context == "" {
		log.Fatal("Error retrieving context from env variables")
	}

	if consulOption.Service == "" {
		log.Fatal("Error retrieving service from env variables")
	}

	if consulOption.Port != "" {
		t.URL += t.Port
	}

	return consulOption.URL + consulOption.Context + consulOption.Service
}

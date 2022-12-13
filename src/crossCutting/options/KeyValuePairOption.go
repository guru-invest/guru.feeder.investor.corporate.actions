package options

import (
	"encoding/base64"

	log "github.com/sirupsen/logrus"
)

type KeyValuePairOption []struct {
	LockIndex   int    `json:"LockIndex"`
	Key         string `json:"Key"`
	Flags       int    `json:"Flags"`
	Value       string `json:"Value"`
	CreateIndex int    `json:"CreateIndex"`
	ModifyIndex int    `json:"ModifyIndex"`
}

func (kv KeyValuePairOption) DecodeString() []byte {
	rawString, err := base64.StdEncoding.DecodeString(kv[0].Value)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error while decoding config")
	}
	return rawString
}

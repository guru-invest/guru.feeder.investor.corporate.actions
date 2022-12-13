package repository

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/options"
	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
)

type WalletConnector struct {
	_baseURL string
}

func NewWalletConnector() WalletConnector {
	return WalletConnector{
		_baseURL: options.OPTIONS.ENDPOINTS.WalletSync,
	}
}

func (t WalletConnector) ResyncAVGOMS() error {
	uri := t._baseURL + "/oms/recalcavg"
	client := http_connector.HttpClient{}

	header := map[string]string{
		"Content-type": "application/json",
	}
	_, err := client.Post(uri, nil, header)
	if err != nil {

		return err
	}
	return nil
}

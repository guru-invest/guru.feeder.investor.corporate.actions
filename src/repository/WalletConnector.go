package repository

import (
	"net/http"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	http_connector "github.com/guru-invest/guru.framework/src/infrastructure/http-connector"
)

type WalletConnector struct {
	_baseURL    string
	_HTTPClient http_connector.HttpClient
}

func NewWalletConnector() WalletConnector {
	return WalletConnector{
		_baseURL: options.OPTIONS.ENDPOINTS.WalletSync,
		_HTTPClient: http_connector.HttpClient{
			Header: http.Header{
				"Content-type": []string{"application/json"},
			},
			Timeout: 500 * time.Second,
		},
	}
}

func (t WalletConnector) ResyncAVGInvestor() error {
	uri := t._baseURL + "/b3/recalc/avg"

	_, err := t._HTTPClient.Patch(uri, nil)
	if err != nil {

		return err
	}
	return nil
}

func (t WalletConnector) ResyncAVGManual() error {
	uri := t._baseURL + "/manual/recalcavg"

	_, err := t._HTTPClient.Post(uri, nil)
	if err != nil {

		return err
	}
	return nil
}

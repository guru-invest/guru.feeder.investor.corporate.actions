package repository

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
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

func (t WalletConnector) ResyncAVGInvestor() error {
	uri := t._baseURL + "/b3/recalc/avg"
	client := http_connector.HttpClient{}

	header := map[string]string{
		"Content-type": "application/json",
	}
	_, err := client.Patch(uri, nil, header)
	if err != nil {

		return err
	}
	return nil
}

func (t WalletConnector) ResyncAVGManual() error {
	uri := t._baseURL + "/manual/recalcavg"
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

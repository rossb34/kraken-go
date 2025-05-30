package kraken

import (
	"encoding/json"
	"io"
	"net/http"
)

type KrakenFuturesHttpClient struct {
	baseURL string
}

func NewKrakenFuturesHttpClient() *KrakenFuturesHttpClient {
	return &KrakenFuturesHttpClient{baseURL: "https://futures.kraken.com"}
}

func (c *KrakenFuturesHttpClient) get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: check response code here.
	// I think I always want a 2XX code

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *KrakenFuturesHttpClient) GetInstruments() ([]map[string]any, error) {
	endpoint := "/derivatives/api/v3/instruments"
	url := c.baseURL + endpoint
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var d KrakenFuturesInstrumentsResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if d.Result == "error" {
		// FIXME: include the array of errors
		return nil, &KrakenError{d.Error}
	}

	return d.Instruments, nil
}

func (c *KrakenFuturesHttpClient) GetInstrumentStatus(symbol string) (KrakenFuturesInstrumentStatus, error) {
	endpoint := "/derivatives/api/v3/instruments/" + symbol + "/status"
	url := c.baseURL + endpoint
	body, err := c.get(url)
	if err != nil {
		return KrakenFuturesInstrumentStatus{}, err
	}

	var d KrakenFuturesInstrumentStatusResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenFuturesInstrumentStatus{}, err
	}

	if d.Result == "error" {
		// FIXME: include the array of errors
		return KrakenFuturesInstrumentStatus{}, &KrakenError{d.Error}
	}

	status := KrakenFuturesInstrumentStatus{
		d.Symbol,
		d.IsExperiencingDislocation,
		d.PriceDislocationDirection,
		d.IsExperiencingExtremeVolatility,
		d.ExtremeVolatilityInitialMarginMultiplier,
	}

	return status, nil
}

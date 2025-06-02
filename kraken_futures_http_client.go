package kraken

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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

func (c *KrakenFuturesHttpClient) GetTicker(symbol string) (KrakenFuturesTickerInfo, error) {
	endpoint := "/derivatives/api/v3/tickers"
	url := c.baseURL + endpoint + "/" + symbol
	body, err := c.get(url)
	if err != nil {
		return KrakenFuturesTickerInfo{}, err
	}

	var d KrakenFuturesTickerBySymbolResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenFuturesTickerInfo{}, err
	}

	if d.Result == "error" {
		return KrakenFuturesTickerInfo{}, &KrakenError{d.Error}
	}

	return d.Ticker, nil
}

func (c *KrakenFuturesHttpClient) GetTradeHistory(symbol string, lastTime string) ([]KrakenFuturesTradeInfo, error) {
	// query params
	params := url.Values{}
	params.Add("symbol", symbol)

	// empty string means the query param for "lastTime" is ignored
	if len(lastTime) > 0 {
		params.Add("lastTime", lastTime)
	}
	queryString := params.Encode()

	endpoint := "/derivatives/api/v3/history"
	url := c.baseURL + endpoint + "?" + queryString
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var d KrakenFuturesTradeHistoryResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if d.Result == "error" {
		return nil, &KrakenError{d.Error}
	}

	return d.History, nil
}

func (c *KrakenFuturesHttpClient) GetOrderBook(symbol string) (KrakenFuturesOrderBook, error) {

	// query params
	params := url.Values{}
	params.Add("symbol", symbol)
	queryString := params.Encode()

	endpoint := "/derivatives/api/v3/orderbook"
	url := c.baseURL + endpoint + "?" + queryString
	body, err := c.get(url)
	if err != nil {
		return KrakenFuturesOrderBook{}, err
	}

	var d KrakenFuturesOrderBookResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenFuturesOrderBook{}, err
	}

	if d.Result == "error" {
		return KrakenFuturesOrderBook{}, &KrakenError{d.Error}
	}

	return d.OrderBook, nil
}

// FIXME: implement
func (c *KrakenFuturesHttpClient) GetHistoricalFundingRates(symbol string) ([]KrakenFuturesFundingRate, error) {
	// query params
	params := url.Values{}
	params.Add("symbol", symbol)
	queryString := params.Encode()

	endpoint := "/derivatives/api/v3/historical-funding-rates"
	url := c.baseURL + endpoint + "?" + queryString
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var d KrakenFuturesFundingRateResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if d.Result == "error" {
		return nil, &KrakenError{d.Error}
	}

	return d.Rates, nil
}

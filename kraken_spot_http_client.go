package kraken

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const DEFAULT_TRADES_CAPACITY int = 10

type KrakenSpotHttpClient struct {
	baseURL string
}

func NewKrakenSpotHttpClient() *KrakenSpotHttpClient {
	return &KrakenSpotHttpClient{baseURL: "https://api.kraken.com/0"}
}

func (c *KrakenSpotHttpClient) get(url string) ([]byte, error) {
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

func (c *KrakenSpotHttpClient) GetServerTime() (KrakenSpotServerTime, error) {
	endpoint := "/public/Time"
	url := c.baseURL + endpoint
	body, err := c.get(url)
	if err != nil {
		return KrakenSpotServerTime{}, err
	}

	// Unmarshall the response body into a struct
	var d KrakenSpotServerTimeResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenSpotServerTime{}, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return KrakenSpotServerTime{}, &KrakenError{d.Error[0]}
	}

	return d.Result, nil
}

func (c *KrakenSpotHttpClient) GetSystemStatus() (KrakenSpotSystemStatus, error) {
	endpoint := "/public/SystemStatus"
	url := c.baseURL + endpoint
	body, err := c.get(url)
	if err != nil {
		return KrakenSpotSystemStatus{}, err
	}

	// Unmarshall the response body into a struct
	var d KrakenSpotSystemStatusResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenSpotSystemStatus{}, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return KrakenSpotSystemStatus{}, &KrakenError{d.Error[0]}
	}

	return d.Result, nil
}

func (c *KrakenSpotHttpClient) GetAssetInfo() (map[string]KrakenSpotAssetInfo, error) {
	// TODO: handle argument for getting specific assets, e.g. (assets []string)
	// the default is to get all assets
	endpoint := "/public/Assets"
	url := c.baseURL + endpoint
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshall the response body into a struct
	var d KrakenSpotAssetInfoResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return nil, &KrakenError{d.Error[0]}
	}

	return d.Result, nil
}

func (c *KrakenSpotHttpClient) GetAssetPairs() (map[string]KrakenSpotAssetPairInfo, error) {
	// TODO: handle argument for getting specific asset pairs, e.g. (pairs []string)
	// the default is to get all assets
	// params := url.Values{}
	// params.Add("pair", pair)

	// queryString := params.Encode()

	endpoint := "/public/AssetPairs"
	url := c.baseURL + endpoint
	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshall the response body into a struct
	var d KrakenSpotAssetPairResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return nil, &KrakenError{d.Error[0]}
	}

	return d.Result, nil
}

func (c *KrakenSpotHttpClient) GetTickerInfo(pair string) (KrakenSpotAssetTickerInfo, error) {
	// query params
	params := url.Values{}
	params.Add("pair", pair)
	queryString := params.Encode()

	endpoint := "/public/Ticker"
	url := c.baseURL + endpoint + "?" + queryString
	body, err := c.get(url)
	if err != nil {
		return KrakenSpotAssetTickerInfo{}, err
	}

	// Unmarshall the response body into a struct
	var d KrakenTickerResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenSpotAssetTickerInfo{}, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return KrakenSpotAssetTickerInfo{}, &KrakenError{d.Error[0]}
	}

	return d.Result[pair], nil
}

func (c *KrakenSpotHttpClient) GetOrderBook(pair string, depth int) (KrakenSpotOrderBook, error) {
	// query params
	params := url.Values{}
	params.Add("pair", pair)
	params.Add("count", strconv.Itoa(depth))
	queryString := params.Encode()

	endpoint := "/public/Depth"
	url := c.baseURL + endpoint + "?" + queryString
	body, err := c.get(url)
	if err != nil {
		return KrakenSpotOrderBook{}, err
	}

	// Unmarshall the response body into a struct
	var d KrakenSpotDepthResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenSpotOrderBook{}, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return KrakenSpotOrderBook{}, &KrakenError{d.Error[0]}
	}

	return d.Result[pair], nil
}

func (c *KrakenSpotHttpClient) GetTrades(pair string, since string, count int) (KrakenSpotTradeInfo, error) {

	// query params
	params := url.Values{}
	params.Add("pair", pair)

	// empty string means the query param for "since" is ignored
	if len(since) > 0 {
		params.Add("since", since)
	}

	// count of 0 means the query param for "count" is ignored
	if count > 0 {
		params.Add("count", strconv.Itoa(count))
	}
	queryString := params.Encode()

	endpoint := "/public/Trades"
	url := c.baseURL + endpoint + "?" + queryString

	body, err := c.get(url)
	if err != nil {
		return KrakenSpotTradeInfo{nil, ""}, err
	}

	// Unmarshall the response body into a struct
	var d KrakenSpotTradeResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return KrakenSpotTradeInfo{nil, ""}, err
	}

	if len(d.Error) > 0 {
		// FIXME: I don't know if len of errors is ever greater than 1.
		// I don't see any examples of that on the kraken website
		return KrakenSpotTradeInfo{nil, ""}, &KrakenError{d.Error[0]}
	}

	out := make([]KrakenSpotTradeEntry, 0, DEFAULT_TRADES_CAPACITY)

	result := d.Result
	last := result["last"].(string)

	trades := result[pair].([]any)
	for _, trades_ := range trades {

		// [<price>, <volume>, <time>, <buy/sell>, <market/limit>, <miscellaneous>, <trade_id>]
		trade_info := trades_.([]any)

		trade_price, err := strconv.ParseFloat(trade_info[0].(string), 64)
		if err != nil {
			return KrakenSpotTradeInfo{nil, ""}, err
		}
		trade_size, err := strconv.ParseFloat(trade_info[1].(string), 64)
		if err != nil {
			return KrakenSpotTradeInfo{nil, ""}, err
		}
		trade := KrakenSpotTradeEntry{
			trade_price,
			trade_size,
			trade_info[2].(float64),
			trade_info[3].(string),
			trade_info[4].(string),
			trade_info[5].(string),
			trade_info[6].(float64),
		}
		out = append(out, trade)
	}

	return KrakenSpotTradeInfo{out, last}, nil
}

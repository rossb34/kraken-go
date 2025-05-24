package kraken

type KrakenError struct {
	Message string
}

func (e *KrakenError) Error() string {
	return e.Message
}

type KrakenSpotResponse struct {
	Error  []string       `json:"error"`
	Result map[string]any `json:"result"`
}

type KrakenSpotServerTime struct {
	UnixTime int64  `json:"unixtime"`
	RFC1123  string `json:"rfc1123"`
}

type KrakenSpotServerTimeResponse struct {
	Error  []string             `json:"error"`
	Result KrakenSpotServerTime `json:"result"`
}

type KrakenSpotSystemStatus struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type KrakenSpotSystemStatusResponse struct {
	Error  []string               `json:"error"`
	Result KrakenSpotSystemStatus `json:"result"`
}

type KrakenSpotAssetInfo struct {
	AssetClass      string  `json:"aclass"`
	AlternateName   string  `json:"altname"`
	Decimals        int     `json:"decimals"`
	DisplayDecimals int     `json:"display_decimals"`
	CollateralValue float64 `json:"collateral_value,omitempty"`
	Status          string  `json:"status"`
	MarginRate      string  `json:"margin_rate,omitempty"`
}

type KrakenSpotAssetInfoResponse struct {
	Error  []string                       `json:"error"`
	Result map[string]KrakenSpotAssetInfo `json:"result"`
}

type KrakenSpotAssetPairInfo struct {
	AlternateName      string  `json:"altname"`
	AssetClassBase     string  `json:"aclass_base"`
	Base               string  `json:"base"`
	AssetClasQuote     string  `json:"aclass_quote"`
	Quote              string  `json:"quote"`
	Lot                string  `json:"lot,omitempty"`
	CostDecimals       int     `json:"cost_decimals"`
	PairDecimals       int     `json:"pair_decimals"`
	LotDecimals        int     `json:"lot_decimals"`
	LotMultiplier      int     `json:"lot_multiplier"`
	LeverageBuy        []int   `json:"leverage_buy"`
	LeverageSell       []int   `json:"leverage_sell"`
	Fees               [][]any `json:"fees"`
	FeesMaker          [][]any `json:"fees_maker"`
	FeeVolumneCurrency string  `json:"fees_volume_currency"`
	MarginCall         int     `json:"margin_call"`
	MarginStop         int     `json:"margin_stop"`
	CostMin            string  `json:"costmin"`
	TickSize           string  `json:"tick_size"`
	Status             string  `json:"status"`
	LongPositionLimit  int     `json:"long_position_limit"`
	ShortPositionLimit int     `json:"short_position_limit"`
}

type KrakenSpotAssetPairResponse struct {
	Error  []string                           `json:"error"`
	Result map[string]KrakenSpotAssetPairInfo `json:"result"`
}

type KrakenSpotAssetTickerInfo struct {
	Ask        []string `json:"a"`
	Bid        []string `json:"b"`
	LastTrade  []string `json:"c"`
	Volume     []string `json:"v"`
	VWAP       []string `json:"p"`
	TradeCount []int    `json:"t"`
	Low        []string `json:"l"`
	High       []string `json:"h"`
	Open       string   `json:"o"`
}

type KrakenTickerResponse struct {
	Error  []string                             `json:"error"`
	Result map[string]KrakenSpotAssetTickerInfo `json:"result"`
}

// type KrakenSpotPriceLevel struct {
// 	data []any
// }

// [<price>, <volume>, <timestamp>]
type KrakenSpotPriceLevel []any

type KrakenSpotOrderBook struct {
	Asks []KrakenSpotPriceLevel `json:"asks"`
	Bids []KrakenSpotPriceLevel `json:"bids"`
}

type KrakenSpotDepthResponse struct {
	Error  []string                       `json:"error"`
	Result map[string]KrakenSpotOrderBook `json:"result"`
}

type KrakenSpotTradeEntry struct {
	Price              float64
	Quantity           float64
	TransactTime       float64
	AggressorSide      string
	AggressorOrderType string
	Misc               string
	TradeId            float64
}

type KrakenSpotTradeInfo struct {
	Trades []KrakenSpotTradeEntry
	Last   string
}

type KrakenSpotTradeResponse struct {
	Error  []string       `json:"error"`
	Result map[string]any `json:"result"`
}

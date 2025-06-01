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

type KrakenFuturesInstrumentsResponse struct {
	Instruments []map[string]any `json:"instruments,omitempty"`
	Result      string           `json:"result"`
	ServerTime  string           `json:"serverTime"`
	Error       string           `json:"error,omitempty"`
	Errors      []string         `json:"errors,omitempty"`
}

type KrakenFuturesInstrumentStatus struct {
	Symbol                                   string
	IsExperiencingDislocation                bool
	PriceDislocationDirection                string
	IsExperiencingExtremeVolatility          bool
	ExtremeVolatilityInitialMarginMultiplier int
}

type KrakenFuturesInstrumentStatusResponse struct {
	Result                                   string   `json:"result"`
	ServerTime                               string   `json:"serverTime"`
	Error                                    string   `json:"error,omitempty"`
	Errors                                   []string `json:"errors,omitempty"`
	Symbol                                   string   `json:"tradeable,omitempty"`
	IsExperiencingDislocation                bool     `json:"experiencingDislocation,omitempty"`
	PriceDislocationDirection                string   `json:"priceDislocationDirection,omitempty"`
	IsExperiencingExtremeVolatility          bool     `json:"experiencingExtremeVolatility,omitempty"`
	ExtremeVolatilityInitialMarginMultiplier int      `json:"extremeVolatilityInitialMarginMultiplier,omitempty"`
}

type KrakenFuturesTickerInfo struct {
	Symbol                   string         `json:"symbol"`
	Last                     float64        `json:"last,omitempty"`
	LastTime                 string         `json:"lastTime,omitempty"`
	LastSize                 float64        `json:"lastSize,omitempty"`
	Tag                      string         `json:"tag,omitempty"`
	Pair                     string         `json:"pair,omitempty"`
	MarkPrice                float64        `json:"markPrice,omitempty"`
	BidPrice                 float64        `json:"bid,omitempty"`
	BidSize                  float64        `json:"bidSize,omitempty"`
	AskPrice                 float64        `json:"ask,omitempty"`
	AskSize                  float64        `json:"askSize,omitempty"`
	Vol24h                   float64        `json:"vol24h,omitempty"`
	VolumeQuote              float64        `json:"volumeQuote,omitempty"`
	OpenInterest             float64        `json:"openInterest,omitempty"`
	Open24h                  float64        `json:"open24h,omitempty"`
	High24h                  float64        `json:"high24h,omitempty"`
	Low24h                   float64        `json:"low24h,omitempty"`
	ExtrinsicValue           float64        `json:"extrinsicValue,omitempty"`
	FundingRate              float64        `json:"fundingRate,omitempty"`
	FundingRatePrediction    float64        `json:"fundingRatePrediction,omitempty"`
	IsSuspended              bool           `json:"suspended,omitempty"`
	IndexPrice               float64        `json:"indexPrice,omitempty"`
	IsPostOnly               bool           `json:"postOnly,omitempty"`
	PercentChange24h         float64        `json:"change24h,omitempty"`
	Greeks                   map[string]any `json:"greeks,omitempty"`
	IsUnderlyingMarketClosed bool           `json:"isUnderlyingMarketClosed,omitempty"`
}

type KrakenFuturesTickerBySymbolResponse struct {
	Ticker     KrakenFuturesTickerInfo `json:"ticker,omitempty"`
	Result     string                  `json:"result"`
	ServerTime string                  `json:"serverTime"`
	Error      string                  `json:"error,omitempty"`
	Errors     []string                `json:"errors,omitempty"`
}

type KrakenFuturesTradeInfo struct {
	TransactTime                  string  `json:"time"`
	Price                         float64 `json:"price"`
	Size                          float64 `json:"size,omitempty"`
	Side                          string  `json:"side,omitempty"`
	TradeId                       int64   `json:"trade_id,omitempty"`
	Type                          string  `json:"type,omitempty"`
	UID                           string  `json:"uid,omitempty"`
	InstrumentIdentificationType  string  `json:"instrument_identification_type,omitempty"`
	ISIN                          string  `json:"isin,omitempty"`
	ExecutionVenue                string  `json:"execution_venue,omitempty"`
	PriceNotation                 string  `json:"price_notation,omitempty"`
	PriceCurrency                 string  `json:"price_currency,omitempty"`
	NotionalAmount                float64 `json:"notional_amount,omitempty"`
	NotionalCurrency              string  `json:"notional_currenct,omitempty"`
	PublicationTime               string  `json:"publication_Type,omitempty"`
	PublicationVenue              string  `json:"publication_venue,omitempty"`
	TransactionIdentificationCode string  `json:"transaction_identification_code,omitempty"`
	IsToBeCleared                 bool    `json:"to_be_cleared,omitempty"`
}

type KrakenFuturesTradeHistoryResponse struct {
	History    []KrakenFuturesTradeInfo `json:"history,omitempty"`
	Result     string                   `json:"result"`
	ServerTime string                   `json:"serverTime"`
	Error      string                   `json:"error,omitempty"`
	Errors     []string                 `json:"errors,omitempty"`
}

type KrakenFuturesOrderBook struct {
	Bids [][]any `json:"bids"`
	Asks [][]any `json:"asks"`
}

type KrakenFuturesOrderBookResponse struct {
	OrderBook  KrakenFuturesOrderBook `json:"orderBook"`
	Result     string                 `json:"result"`
	ServerTime string                 `json:"serverTime"`
	Error      string                 `json:"error,omitempty"`
	Errors     []string               `json:"errors,omitempty"`
}

type KrakenFuturesFundingRate struct {
	FundingRate         float64 `json:"fundingRate"`
	RelativeFundingRate float64 `json:"relativeFundingRate"`
	Timestamp           string  `json:"timestamp"`
}

type KrakenFuturesFundingRateResponse struct {
	Rates      []KrakenFuturesFundingRate `json:"rates,omitempty"`
	Result     string                     `json:"result"`
	ServerTime string                     `json:"serverTime"`
	Error      string                     `json:"error,omitempty"`
	Errors     []string                   `json:"errors,omitempty"`
}

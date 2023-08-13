package models

type YahooChart struct {
	Chart ChartData `json:"chart"`
}

type ChartData struct {
	Result      []ChartResult `json:"result"`
	ErrorStatus error         `json:"error"`
}

type ChartResult struct {
	Meta       ChartResultMeta       `json:"meta"`
	Timestamp  ChartResultTimestamp  `json:"timestamp"`
	Indicators ChartResultIndicators `json:"indicators"`
}

type ChartResultMeta struct {
	Currency             string   `json:"currency"`
	Symbol               string   `json:"symbol"`
	ExchangeName         string   `json:"exhangeName"`
	InstrumentType       string   `json:"instrumentType"`
	FirstTradeDate       string   `json:"firstTradeDate"`
	RegularMarketTime    string   `json:"regularMarketTime"`
	GmtOffset            string   `json:"gmtoffset"`
	Timezone             string   `json:"timezone"`
	ExchangeTimezoneName string   `json:"exchangeTimezoneName"`
	RegularMarketPrice   float32  `json:"regularMarketPrice"`
	ChartPreviousClose   float32  `json:"chartPreviousClose"`
	PriceHint            float32  `json:"priceHint"`
	DataGranularity      string   `json:"dataGranularity"`
	TotalRange           string   `json:"range"`
	ValidRanges          []string `json:"validRanges"`
}

type ChartResultTimestamp struct {
	Timestamp []int64 `json:"timestamp"`
}

type ChartResultIndicators struct {
	Quote        []ChartResultIndicatorsQuote    `json:"quote"`
	AdjCloseData []ChartResultIndicatorsAdjClose `json:"adjclose"`
}

type ChartResultIndicatorsQuote struct {
	High   []float32 `json:"high"`
	Volume []float32 `json:"volume"`
	Close  []float32 `json:"close"`
	Low    []float32 `json:"low"`
	Open   []float32 `json:"open"`
}

type ChartResultIndicatorsAdjClose struct {
	AdjClose []float32 `json:"adjclose"`
}

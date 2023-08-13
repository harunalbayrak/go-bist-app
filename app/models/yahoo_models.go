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

type YahooQuoteResponse struct {
	QuoteResponse QuoteResponseData `json:"quoteResponse"`
}

type QuoteResponseData struct {
	Result      []QuoteResponseResult `json:"result"`
	ErrorStatus error                 `json:"error"`
}

type QuoteResponseResult struct {
	Language                          string  `json:"result"`
	Region                            string  `json:"region"`
	QuoteType                         string  `json:"quoteType"`
	TypeDisp                          string  `json:"typeDisp"`
	QuoteSourceName                   string  `json:"quoteSourceName"`
	Triggerable                       bool    `json:"triggerable"`
	RegularMarketChange               float32 `json:"regularMarketChange"`
	RegularMarketTime                 int64   `json:"regularMarketTime"`
	RegularMarketDayHigh              float32 `json:"regularMarketDayHigh"`
	RegularMarketDayRange             string  `json:"regularMarketDayRange"`
	RegularMarketDayLow               float32 `json:"regularMarketDayLow"`
	RegularMarketVolume               int64   `json:"regularMarketVolume"`
	RegularMarketPreviousClose        float32 `json:"regularMarketPreviousClose"`
	Bid                               float32 `json:"bid"`
	Ask                               float32 `json:"ask"`
	BidSize                           float32 `json:"bidSize"`
	AskSize                           float32 `json:"askSize"`
	FullExchangeName                  string  `json:"fullExchangeName"`
	FinancialCurrency                 string  `json:"financialCurrency"`
	RegularMarketOpen                 float32 `json:"regularMarketOpen"`
	AverageDailyVolume3Month          int64   `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day           int64   `json:"averageDailyVolume10Day"`
	FiftyTwoWeekLowChange             float32 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float32 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekHighChange            float32 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float32 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float32 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh                  float32 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekChangePercent         float32 `json:"fiftyTwoWeekChangePercent"`
	Currency                          string  `json:"currency"`
	RegularMarketPrice                float32 `json:"regularMarketPrice"`
	RegularMarketChangePercent        float32 `json:"regularMarketChangePercent"`
	Exchange                          string  `json:"exchange"`
	ShortName                         string  `json:"shortName"`
	LongName                          string  `json:"longName"`
	MessageBoardId                    string  `json:"messageBoardId"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int64   `json:"gmtOffSetMilliseconds"`
	Market                            string  `json:"market"`
	EsgPopulated                      bool    `json:"esgPopulated"`
	Tradeable                         bool    `json:"tradeable"`
	CryptoTradeable                   bool    `json:"cryptoTradeable"`
	MarketState                       string  `json:"marketState"`
	EarningsTimestampStart            int64   `json:"earningsTimestampStart"`
	EarningsTimestampEnd              int64   `json:"earningsTimestampEnd"`
	TrailingAnnualDividendRate        float32 `json:"trailingAnnualDividendRate"`
	TrailingAnnualDividendYield       float32 `json:"trailingAnnualDividendYield"`
	EpsTrailingTwelveMonths           float32 `json:"epsTrailingTwelveMonths"`
	SharesOutstanding                 int64   `json:"sharesOutstanding"`
	BookValue                         float32 `json:"bookValue"`
	FiftyDayAverage                   float32 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float32 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float32 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float32 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float32 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float32 `json:"twoHundredDayAverageChangePercent"`
	MarketCap                         int64   `json:"marketCap"`
	PriceToBook                       float32 `json:"priceToBook"`
	SourceInterval                    float32 `json:"sourceInterval"`
	ExchangeDataDelayedBy             float32 `json:"exchangeDataDelayedBy"`
	FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
	PriceHint                         float32 `json:"priceHint"`
	Symbol                            string  `json:"symbol"`
}

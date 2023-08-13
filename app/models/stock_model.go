package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/harunalbayrak/go-finance/pkg/yahoo"
)

type Stock struct {
	Code string `json:"code"`
}

func (stock *Stock) GetYahooChart(interval string, totalRange string) (*YahooChart, error) {
	baseURL := "https://query1.finance.yahoo.com/v8/finance"

	requestURL := fmt.Sprintf("%s/chart/%s.is?metrics=high?&interval=%s&range=%s", baseURL, stock.Code, interval, totalRange)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var yahooModel YahooChart
	json.Unmarshal(resBody, &yahooModel)

	return &yahooModel, err
}

func (stock *Stock) GetYahooQuoteResponse(cookie *http.Cookie, crumb string) (*YahooQuoteResponse, error) {
	baseURL := "https://query2.finance.yahoo.com/v7/finance"
	userAgentKey := "User-Agent"
	userAgentValue := "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

	requestURL := fmt.Sprintf("%s/quote?symbols=%s.is&crumb=%s", baseURL, stock.Code, crumb)

	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(userAgentKey, userAgentValue)
	req.AddCookie(&http.Cookie{
		Name: cookie.Name, Value: cookie.Value, MaxAge: 60,
	})
	req.Header.Set("Accept", "application/json")

	resBody, err := yahoo.GetRequestBody(req)

	var yahooQuoteResponse YahooQuoteResponse
	json.Unmarshal(resBody, &yahooQuoteResponse)

	return &yahooQuoteResponse, err
}

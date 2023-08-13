package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	// fmt.Printf("client: response body: %s\n", resBody)

	var yahooModel YahooChart
	json.Unmarshal(resBody, &yahooModel)

	// fmt.Println(yahooModel)

	return &yahooModel, nil
}

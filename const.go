package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	api            string = "https://query2.finance.yahoo.com/v7/finance/"
	chartApi       string = "https://query1.finance.yahoo.com/v8/finance/chart/"
	cookieLink     string = "https://fc.yahoo.com/"
	crumbLink      string = "https://query1.finance.yahoo.com/v1/test/getcrumb"
	userAgentKey   string = "User-Agent"
	userAgentValue string = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
)

func getRequest(req *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return b
}

func getCookie() *http.Cookie {
	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", cookieLink, nil)
	req.Header.Set(userAgentKey, userAgentValue)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie: %s=%s\n", cookie.Name, cookie.Value)
	}

	return cookies[0]
}

func getCrumb(cookie *http.Cookie) string {
	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", "https://query1.finance.yahoo.com/v1/test/getcrumb", nil)
	req.Header.Set(userAgentKey, userAgentValue)
	req.AddCookie(&http.Cookie{
		Name: cookie.Name, Value: cookie.Value, MaxAge: 60,
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}

func deneme1() {
	symbol := "EKIZ.IS"

	cookie := getCookie()
	crumb := getCrumb(cookie)

	requestLink := api + "quote?symbols=" + symbol + "&crumb=" + crumb

	fmt.Printf("Crumb: %s\n", crumb)
	fmt.Printf("RequestLink: %s\n", requestLink)

	req, err := http.NewRequest("GET", requestLink, nil)
	req.Header.Set(userAgentKey, userAgentValue)
	req.AddCookie(&http.Cookie{
		Name: cookie.Name, Value: cookie.Value, MaxAge: 60,
	})
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	b := getRequest(req)

	fmt.Println(string(b))
}

func getChart(symbol string, interval string, totalRange string) {
	requestLink := chartApi + symbol + "?interval=" + interval + "&range=" + totalRange

	req, _ := http.NewRequest("GET", requestLink, nil)
	req.Header.Set("Accept", "application/json")

	b := getRequest(req)

	fmt.Println(string(b))
}

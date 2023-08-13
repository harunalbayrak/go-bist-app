package yahoo

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

func GetRequestBody(req *http.Request) ([]byte, error) {
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

	return b, err
}

func GetCookie() (*http.Cookie, error) {
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

	return cookies[0], err
}

func GetCrumb(cookie *http.Cookie) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", crumbLink, nil)
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

	fmt.Printf("Crumb: %s\n", string(b))

	return string(b), err
}

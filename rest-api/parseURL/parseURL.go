package parseURL

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	parsedURL *url.URL
	err       error
)

func ParseURL(requestURL, apiKey string) ([]byte, error) {
	parsedURL, err = url.ParseRequestURI(requestURL)
	if err != nil {
		flag.Usage()
		return nil, fmt.Errorf("validation error: URL is not valid: %s", err)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // Ignore certificates

	client := &http.Client{}
	request, err := http.NewRequest("GET", parsedURL.String(), nil)
	request.Header.Set("X-Auth-Apikey", apiKey)
	response, _ := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http request error %s", err)
	}

	defer response.Body.Close()
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll error %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("invalid output (HTTP code %d), %s", response.StatusCode, respBody)
	}

	return respBody, nil

}

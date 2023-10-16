package getResponseBody

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetResponseBody(parsedURL *url.URL, token string) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // Ignore certificates
	client := &http.Client{}
	request, err := http.NewRequest("GET", parsedURL.String(), nil)
	request.Header.Set("X-Auth-Token", token)
	request.Header.Set("Content-Type", "application/json")
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

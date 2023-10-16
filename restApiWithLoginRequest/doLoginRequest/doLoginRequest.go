package doLoginRequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Token struct {
	TokenString string `json:"token"`
}

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type LoginResponse struct {
	Token Token `json:"token"`
}

func GetToken(parsedURL, username, password string) (string, error) {

	loginRequest := map[string]interface{}{
		"password": map[string]string{
			"username": username,
			"password": password,
		},
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // Ignore certificates

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("marshal error : %s", err)
	}

	response, err := http.Post(parsedURL+"/v1/auth", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("http auth post error : %s", err)
	}

	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("readAll error: %s", err)
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("invalid output (HTTP Code %d): %s)", response.StatusCode, string(body))
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(res, &loginResponse)
	if err != nil {
		return "", fmt.Errorf("unmarshall error: %s", err)
	}
	return loginResponse.Token.TokenString, nil

}

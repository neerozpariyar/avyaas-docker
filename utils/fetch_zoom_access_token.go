package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// FetchZoomAccessToken fetches the access token from the Zoom API using credentials from the configuration.
func FetchZoomAccessToken(email string) (string, error) {
	// Determine which Zoom account to use based on the email
	var zoomAccount string
	switch email {
	case "abc@gmail.com":
		zoomAccount = "zoom1"
	case "xyz@gmail.com":
		zoomAccount = "zoom2"
	default:
		return "", fmt.Errorf("no Zoom account found for email: %s", email)
	}

	// Load credentials for the selected Zoom account
	creds := viper.GetStringMapString(zoomAccount)
	if creds == nil {
		return "", fmt.Errorf("credentials not found for Zoom account: %s", zoomAccount)
	}

	// Construct the request body to fetch access token
	requestBody := map[string]string{
		"grant_type": "account_credentials",
		"account_id": creds["account_id"],
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Send request to fetch access token
	tokenURL := fmt.Sprintf("https://zoom.us/oauth/token?grant_type=account_credentials&account_id=%s", creds["account_id"])
	request, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}
	request.SetBasicAuth(creds["client_id"], creds["client_secret"])
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch access token. Status code: %d", response.StatusCode)
	}

	// Decode the response to get the access token
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(response.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

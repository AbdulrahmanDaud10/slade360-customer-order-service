package app

import (
	"net/http"
	"os"
	"time"
)

var (
	BaseLiveEndpoiint    = os.Getenv("BASE_LIVE_ENDPOINT")
	BaseSandBoxEndpoiint = os.Getenv("BASE_SANDBOX_ENDPOINT")
	ApiKey               = os.Getenv("API_KEY")
	UserName             = os.Getenv("USERNAME")
)

type AtClient struct {
	ApiKey     string
	Endpoint   string
	Username   string
	HttpClient *http.Client
}

// New returns an instance of an Africa's Talking client reusbale across different products.
func NewSMS(sandbox bool) *AtClient {
	AtClient := &AtClient{
		ApiKey:     ApiKey,
		Endpoint:   BaseLiveEndpoiint,
		HttpClient: &http.Client{Timeout: time.Second * 5},
		Username:   UserName,
	}

	if sandbox {
		AtClient.Endpoint = BaseLiveEndpoiint
	} else {
		AtClient.Endpoint = BaseSandBoxEndpoiint
	}

	return AtClient
}

// SetHTTPClient can be used to override the default client with a custom set one.
func (at *AtClient) SetHTTPClient(httpClient *http.Client) *AtClient {
	at.HttpClient = httpClient

	return at
}

// setDefaultHeaders sets the standard headers required by the Africa's Talking API.
func (at *AtClient) SetDefaultHeaders(request *http.Request) *http.Request {
	request.Header.Set("apiKey", at.ApiKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request
}

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	smsPath = "messaging"
)

type AtClient struct {
	ApiKey     string
	Endpoint   string
	Username   string
	HttpClient *http.Client
}

type (
	// BulkSMSInput is passed to SendBulkSMS as a parameter.
	BulkSMSInput struct {
		To      []string
		Message string
		From    string
	}

	// BulkSMSRecipient is returned as part of the BulkSMSResponse.
	BulkSMSRquest struct {
		StatusCode uint   `json:"statusCode"`
		Number     string `json:"number"`
		Status     string `json:"status"`
		Cost       string `json:"cost"`
		MessageID  string `json:"messageId"`
	}

	// BulkSMSResponse is returned by SendBulkSMS as a response.
	BulkSMSResponse struct {
		SMSMessageData struct {
			Message    string          `json:"Message"`
			Recipients []BulkSMSRquest `json:"Recipients"`
		} `json:"SMSMessageData"`
	}
)

// SendBulkSMS makes a POST request to send bulk SMS's the Africa's Talking and returns a response.
// It uses opinionated defaults.
func (at *AtClient) SendBulkSMS(ctx context.Context, input BulkSMSInput) (BulkSMSResponse, error) {
	bulkSMSResponse := BulkSMSResponse{}

	form := url.Values{
		"username":             {at.Username},
		"to":                   {strings.Join(input.To, ",")},
		"message":              {input.Message},
		"bulkSMSMode":          {"1"},
		"enqueue":              {"1"},
		"retryDurationInHours": {"1"},
	}

	if input.From != "" {
		form.Set("from", input.From)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, at.Endpoint+smsPath, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return bulkSMSResponse, err
	}
	// req = at.setDefaultHeaders(req)

	resp, err := at.HttpClient.Do(req)
	if err != nil {
		return bulkSMSResponse, err
	}

	respData, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return bulkSMSResponse, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return bulkSMSResponse, fmt.Errorf("status code: %s: %q", resp.Status, respData)
	}

	if err := json.Unmarshal(respData, &bulkSMSResponse); err != nil {
		return bulkSMSResponse, err
	}

	return bulkSMSResponse, nil
}

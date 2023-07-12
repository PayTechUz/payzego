package payzego

import "net/http"

type PayzeClient struct {
	baseURL   string
	apiKey    string
	apiSecret string
	client    *http.Client
}

func NewPayzeClient(baseURL, apiKey, apiSecret string) *PayzeClient {
	return &PayzeClient{
		baseURL:   baseURL,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		client:    http.DefaultClient,
	}
}

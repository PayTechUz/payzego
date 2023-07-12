package payzego

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *PayzeClient) PayWithCard(ctx context.Context, amount float64, currency, cardToken, hookURL string, hookRefund bool) (*PayWithCardResponse, error) {
	requestBody := PayWithCardRequest{
		Method:    "payWithCard",
		APIKey:    c.apiKey,
		APISecret: c.apiSecret,
		Data: PayWithCardData{
			Amount:     amount,
			Currency:   currency,
			CardToken:  cardToken,
			HookURL:    hookURL,
			HookRefund: hookRefund,
		},
	}
	return c.do(ctx, &requestBody, http.MethodPost)
}

func (c *PayzeClient) do(ctx context.Context, requestBody interface{}, method string) (*PayWithCardResponse, error) {
	var payWithCardResp PayWithCardResponse
	requestBytes, err := json.Marshal(requestBody)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL, bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &payWithCardResp)
	if err != nil {
		return nil, err
	}

	if payWithCardResp.Response.Error != "" {
		errMsg := fmt.Sprintf("error was occured with payWithCard: %s", payWithCardResp.Response.Error)
		return nil, errors.New(errMsg)
	}

	if payWithCardResp.Message != "" && payWithCardResp.DateTime != "" {
		errMsg := fmt.Sprintf("error was occured with payWithCard: %s", payWithCardResp.Message)

		return nil, errors.New(errMsg)
	}

	return &payWithCardResp, nil
}

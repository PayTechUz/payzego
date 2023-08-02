package payzego

import (
	"context"
	"net/http"
)

func (c *PayzeClient) PayWithCard(ctx context.Context, amount float64, currency, cardToken, hookURL string, hookRefund bool) (respBody *PayWithCardResponse, err error) {
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

	var responseData PayWithCardResponse

	err = c.doRequest(ctx, &requestBody, &responseData, http.MethodPost)

	if err != nil {
		return nil, err
	}
	return &responseData, nil
}

package payzego

import (
	"context"
	"net/http"
)

func (c *PayzeClient) GetTransactionInfo(ctx context.Context, transactionID string) (respBody *GetTransactionInfoResponse, err error) {
	requestBody := GetTransactionInfoRequest{
		Method:    "getTransactionInfo",
		APIKey:    c.apiKey,
		APISecret: c.apiSecret,
		Data: GetTransactionInfoData{
			TransactionID: transactionID,
		},
	}

	var responseData GetTransactionInfoResponse
	err = c.doRequest(ctx, &requestBody, &responseData, http.MethodPost)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

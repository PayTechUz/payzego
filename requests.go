package payzego

type PayWithCardData struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	CardToken  string  `json:"cardToken"`
	HookURL    string  `json:"hookUrl"`
	HookRefund bool    `json:"hookRefund"`
}

type PayWithCardRequest struct {
	Method    string          `json:"method"`
	APIKey    string          `json:"apiKey"`
	APISecret string          `json:"apiSecret"`
	Data      PayWithCardData `json:"data"`
}

type GetTransactionInfoData struct {
	TransactionID string `json:"transactionId"`
}

type GetTransactionInfoRequest struct {
	Method    string                 `json:"method"`
	APIKey    string                 `json:"apiKey"`
	APISecret string                 `json:"apiSecret"`
	Data      GetTransactionInfoData `json:"data"`
}

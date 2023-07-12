package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Muhammadali-Akbarov/payzego"
)

func main() {
	ctx := context.Background()

	client := payzego.NewPayzeClient("https://payze.io/api/v1", "api-key", "api-secret")

	amount := 1.0
	currency := "USD"
	cardToken := "card-token"
	hookURL := "https://your-gateway/v1/success/"
	hookRefund := true

	response, err := client.PayWithCard(ctx, amount, currency, cardToken, hookURL, hookRefund)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response", response)
}

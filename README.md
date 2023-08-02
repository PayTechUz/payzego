# Payze Implementation

This MVP project helps for implementing <a href="https://docs.payze.io/reference/getting-started">payze-doc</a>.

# Installation
1 - `go get -v github.com/paytechuz/payzego`

## Example
```go

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/paytechuz/payzego"
)

func main() {
    ctx := context.Background()

	client := payzego.NewPayzeClient("https://payze.io/api/v1", "api-key","api-secret")

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
```

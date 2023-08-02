package payzego

type PayWithCardResponse struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	Action      string `json:"action"`
	CreatedDate string `json:"createdDate"`
	Response    struct {
		Error           string `json:"error,omitempty"`
		TransactionID   string `json:"transactionId"`
		TransactionInfo struct {
			CanBeCommitted bool    `json:"canBeCommitted"`
			Type           string  `json:"type"`
			FinalAmount    float64 `json:"finalAmount"`
			AmountRefunded float64 `json:"amountRefunded"`
			Currency       string  `json:"currency"`
			Commission     float64 `json:"commission"`
			Refundable     bool    `json:"refundable"`
			Refunded       float64 `json:"refunded"`
			HasSplit       bool    `json:"hasSplit"`
			CreateDate     string  `json:"createDate"`
			Amount         float64 `json:"amount"`
			Sandbox        bool    `json:"sandbox"`
			TransactionID  string  `json:"transactionId"`
			CommitDate     string  `json:"commitDate"`
			Status         string  `json:"status"`
		} `json:"transactionInfo"`
	} `json:"response"`
	Request struct {
		Amount     float64 `json:"amount"`
		HookRefund bool    `json:"hookRefund"`
		Currency   string  `json:"currency"`
		HookURL    string  `json:"hookUrl"`
		CardToken  string  `json:"cardToken"`
	} `json:"request"`
	Headers struct {
		ContentLength   string `json:"content-length"`
		CDNLoop         string `json:"cdn-loop"`
		CFIPCountry     string `json:"cf-ipcountry"`
		CFRay           string `json:"cf-ray"`
		XForwardedProto string `json:"x-forwarded-proto"`
		XForwardedPort  string `json:"x-forwarded-port"`
		XForwardedFor   string `json:"x-forwarded-for"`
		XRealIP         string `json:"x-real-ip"`
		XForwardedHost  string `json:"x-forwarded-host"`
		CFVisitor       string `json:"cf-visitor"`
		Host            string `json:"host"`
		Connection      string `json:"connection"`
		ContentType     string `json:"content-type"`
		CFConnectingIP  string `json:"cf-connecting-ip"`
		AcceptEncoding  string `json:"accept-encoding"`
		UserAgent       string `json:"user-agent"`
	} `json:"headers"`
}

type GetTransactionInfoResponse struct {
	Response struct {
		Amount         float64  `json:"amount"`
		Log            []string `json:"log"`
		Sandbox        bool     `json:"sandbox"`
		CanBeCommitted bool     `json:"canBeCommitted"`
		Type           string   `json:"type"`
		TransactionID  string   `json:"transactionId"`
		CommitDate     string   `json:"commitDate"`
		FinalAmount    float64  `json:"finalAmount"`
		AmountRefunded float64  `json:"amountRefunded"`
		Currency       string   `json:"currency"`
		Commission     float64  `json:"commission"`
		Refundable     bool     `json:"refundable"`
		Refunded       float64  `json:"refunded"`
		HasSplit       bool     `json:"hasSplit"`
		CreateDate     string   `json:"createDate"`
		Status         string   `json:"status"`
	} `json:"response"`
}

type ErrorResponse struct {
	Message  string `json:"message"`
	DateTime string `json:"dateTime"`
}

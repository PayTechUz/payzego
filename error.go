package payzego

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *PayzeClient) hasError(r []byte, reqBody interface{}) (*ErrorResponse, error) {
	var errResponse *ErrorResponse
	err := json.Unmarshal(r, &errResponse)
	if err != nil {
		return nil, err
	}

	if errResponse.Message != "" && errResponse.DateTime != "" {
		errMsg := fmt.Sprintf("error occurred with payze: %s, body: %s", errResponse.Message, reqBody)
		return nil, errors.New(errMsg)
	}
	return nil, nil
}

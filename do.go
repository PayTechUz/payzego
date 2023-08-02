package payzego

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (c *PayzeClient) doRequest(ctx context.Context, requestBody, responseBody interface{}, method string) error {
	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL, bytes.NewBuffer(requestBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for errors in the response
	errResponse, err := c.hasError(respBody, requestBody)
	if err != nil {
		return err
	}

	// If there is an error response, return it
	if errResponse != nil {
		return errors.New("error occurred with payze")
	}

	// Unmarshal the response into the provided responseBody struct
	if err := json.Unmarshal(respBody, responseBody); err != nil {
		return err
	}

	return nil
}

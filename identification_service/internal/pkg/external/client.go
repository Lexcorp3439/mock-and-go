package external

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"heisenbug/identification/internal/pkg/model"
	"io"
	"net/http"
	"os"
	"time"
)

// Client is an interface for calling externak.Client.
type Client interface {
	Identification(ctx context.Context, phone string) (*model.IdentificationResponse, error)
}

type client struct {
	httpClient *http.Client
}

// NewClient .
func NewClient() Client {
	tran := &http.Transport{
		MaxIdleConns:      100,
		IdleConnTimeout:   90 * time.Second,
		DisableKeepAlives: true,
	}
	roundTran := http.RoundTripper(tran)
	timeout := 5 * time.Second
	return &client{
		httpClient: &http.Client{
			Transport: roundTran,
			Timeout:   timeout,
		},
	}
}

// Identification - .
func (c *client) Identification(ctx context.Context, phone string) (*model.IdentificationResponse, error) {
	api := os.Getenv("EXTERNAL_API")
	uri := fmt.Sprintf("%s/%s", api, "identification") // #nosec
	fmt.Println(uri)

	body := model.IdentificationRequest{
		Phone: phone,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		uri,
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return nil, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.StatusCode)
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("not 200 status code %d", response.Status))
	}
	defer response.Body.Close()
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result model.IdentificationResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

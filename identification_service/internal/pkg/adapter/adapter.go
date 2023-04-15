package adapter

import (
	"encoding/json"
	"heisenbug/identification/internal/pkg/model"
	"io"
	"net/http"
)

func ResponseToIdentificationModel(resp *http.Response) (*model.IdentificationResponse, error) {
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
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

package adapters

import (
	"context"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
)

const CORRENTLY_BASE_URL = "https://api.corrently.io/v2.0"

func (a Adapter) GetGSIPrediction(ctx context.Context, zipcode string) (*GSIResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, CORRENTLY_BASE_URL+"/gsi/prediction?q="+zipcode+"&account="+a.correntlyAPIKey, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Green-Backend")
	response, err := a.doRequest(req)
	if err != nil {
		return nil, err
	}

	var gsiResponse *GSIResponse
	err = sonic.Unmarshal(response, &gsiResponse)
	if err != nil {
		return nil, err
	}

	return gsiResponse, nil
}

func (a Adapter) doRequest(req *http.Request) ([]byte, error) {
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func(adapter Adapter) {
		err := resp.Body.Close()
		if err != nil {
			adapter.logger.Error(err.Error())
		}
	}(a)

	return b, nil
}

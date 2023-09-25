package adapters

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bytedance/sonic"
)

const CORRENTLY_BASE_URL = "https://api.corrently.io/v2.0"

func (a Adapter) GetGSIPrediction(ctx context.Context, zipcode string) (*GSIResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, CORRENTLY_BASE_URL+"/gsi/prediction?zip="+zipcode+"&account="+a.correntlyAPIKey, nil)
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

	fmt.Println(gsiResponse)

	return gsiResponse, nil
}

func (a Adapter) GetLocalPricePrediction(ctx context.Context, zipcode string) (*LocalMarketpriceResponse, error) {
	//TODO Account not set (Request over Rapiddata?: https://rapidapi.com/stromdao-stromdao-default/api/marktdaten-deutschland/)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, CORRENTLY_BASE_URL+"/gsi/marketdata?zip="+zipcode, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Green-Backend")
	response, err := a.doRequest(req)
	if err != nil {
		return nil, err
	}
	var localMarketpriceResponse *LocalMarketpriceResponse
	err = sonic.Unmarshal(response, &localMarketpriceResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(localMarketpriceResponse)

	return localMarketpriceResponse, nil
}

func (a Adapter) GetBestHourForEnergyConsumption(ctx context.Context, zipcode string, numberOfHours string) (*BestHourForEnergyConsumptionResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, CORRENTLY_BASE_URL+"/gsi/besthour?zip="+zipcode+"&timeframe="+numberOfHours+"&account="+a.correntlyAPIKey, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Green-Backend")
	response, err := a.doRequest(req)
	if err != nil {
		return nil, err
	}

	bestHour, err := strconv.ParseBool(string(response))
	if err != nil {
		return nil, err
	}
	fmt.Println(bestHour)

	return &BestHourForEnergyConsumptionResponse{bestHour: bestHour}, nil
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

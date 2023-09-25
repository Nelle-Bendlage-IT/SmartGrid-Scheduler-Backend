package adapters

type GSIResponse struct {
	Support       string     `json:"support"`
	License       string     `json:"license"`
	Info          string     `json:"info"`
	Documentation string     `json:"documentation"`
	Commercial    string     `json:"commercial"`
	Signee        string     `json:"signee"`
	Forecast      []Forecast `json:"forecast"`
	Location      Location   `json:"location"`
}
type Timeframe struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
type Forecast struct {
	Epochtime     int       `json:"epochtime"`
	Eevalue       int       `json:"eevalue"`
	Ewind         int       `json:"ewind"`
	Esolar        int       `json:"esolar"`
	Ensolar       int       `json:"ensolar"`
	Enwind        int       `json:"enwind"`
	Sci           int       `json:"sci"`
	Gsi           float64   `json:"gsi"`
	TimeStamp     int64     `json:"timeStamp"`
	Energyprice   string    `json:"energyprice"`
	Co2GStandard  int       `json:"co2_g_standard"`
	Co2GOekostrom int       `json:"co2_g_oekostrom"`
	Timeframe     Timeframe `json:"timeframe"`
	Iat           int64     `json:"iat"`
	Zip           string    `json:"zip"`
	Signature     string    `json:"signature"`
}
type Location struct {
	Zip       string `json:"zip"`
	City      string `json:"city"`
	Signature string `json:"signature"`
}

type LocalMarketpriceResponse struct {
	Object        string                 `json:"object"`
	URL           string                 `json:"url"`
	License       string                 `json:"license"`
	Documentation string                 `json:"documentation"`
	Support       string                 `json:"support"`
	Data          []LocalMarketpriceData `json:"data"`
}

type LocalMarketpriceData struct {
	StartTimestamp int64   `json:"start_timestamp"`
	EndTimestamp   int64   `json:"end_timestamp"`
	Marketprice    float64 `json:"marketprice"`
	Unit           string  `json:"unit"`
	Localprice     float64 `json:"localprice"`
	Localcell      string  `json:"localcell"`
}

type BestHourForEnergyConsumptionResponse struct {
	bestHour bool
}

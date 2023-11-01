package gsiService

import (
	"time"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapToGSIPredictionResponse(prediciton []*gsi_prediction.GSIPrediction) (*gsi_prediction.GetGSIPredictionResponse, error) {

	var mappedPrediction []*gsi_prediction.GSIPrediction
	for v := range prediciton {
		element := prediciton[v]
		end := timestamppb.New(time.Unix(element.EndTimestamp.Seconds, int64(element.EndTimestamp.Nanos)))
		start := timestamppb.New(time.Unix(element.StartTimestamp.Seconds, int64(element.StartTimestamp.Nanos)))
		// Create GSIPrediction object
		prediction := gsi_prediction.GSIPrediction{
			StartTimestamp: start,
			ZipCode:        uint32(element.ZipCode),
			EndTimestamp:   end,
			Wind:           element.Wind,
			Solar:          element.Solar,
			Gsi:            element.Gsi,
			Co2GStandard:   element.Co2GStandard,
			Co2GOekostrom:  element.Co2GOekostrom,
			Sci:            element.Sci,
			Energyprice:    element.Energyprice,
		}
		//pred := gsi_prediction.GSIPrediction{start, int32(element.ZipCode), end, element.Wind, element.Solar, element.Gsi, uint32(element.Co2GStandard), uint32(element.Co2GOekostrom), element.Sci, element.Energyprice}
		mappedPrediction = append(mappedPrediction, &prediction)
	}

	mappedResp := gsi_prediction.GetGSIPredictionResponse{GsiPredictions: mappedPrediction}
	return &mappedResp, nil

}

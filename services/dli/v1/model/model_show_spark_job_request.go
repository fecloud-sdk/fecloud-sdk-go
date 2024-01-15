package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSparkJobRequest struct {
	BatchId string `json:"batch_id"`
}

func (o ShowSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobRequest", string(data)}, " ")
}

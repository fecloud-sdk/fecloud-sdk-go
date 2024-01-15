package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSparkJobStatusRequest struct {
	BatchId string `json:"batch_id"`
}

func (o ShowSparkJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobStatusRequest", string(data)}, " ")
}

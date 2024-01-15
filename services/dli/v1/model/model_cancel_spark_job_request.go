package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelSparkJobRequest struct {
	BatchId string `json:"batch_id"`
}

func (o CancelSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSparkJobRequest struct{}"
	}

	return strings.Join([]string{"CancelSparkJobRequest", string(data)}, " ")
}

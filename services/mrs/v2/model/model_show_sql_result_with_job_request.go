package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlResultWithJobRequest struct {
	JobExecutionId string `json:"job_execution_id"`

	ClusterId string `json:"cluster_id"`
}

func (o ShowSqlResultWithJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlResultWithJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlResultWithJobRequest", string(data)}, " ")
}

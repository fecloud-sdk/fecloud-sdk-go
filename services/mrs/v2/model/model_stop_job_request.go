package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StopJobRequest struct {
	JobExecutionId string `json:"job_execution_id"`

	ClusterId string `json:"cluster_id"`
}

func (o StopJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobRequest struct{}"
	}

	return strings.Join([]string{"StopJobRequest", string(data)}, " ")
}

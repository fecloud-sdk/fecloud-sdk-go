package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSingleJobExeRequest struct {
	JobExecutionId string `json:"job_execution_id"`

	ClusterId string `json:"cluster_id"`
}

func (o ShowSingleJobExeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleJobExeRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleJobExeRequest", string(data)}, " ")
}

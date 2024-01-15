package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateExecuteJobRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *JobExecution `json:"body,omitempty"`
}

func (o CreateExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"CreateExecuteJobRequest", string(data)}, " ")
}

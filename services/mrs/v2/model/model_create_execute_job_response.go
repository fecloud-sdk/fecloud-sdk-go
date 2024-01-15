package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateExecuteJobResponse struct {
	JobSubmitResult *JobSubmitResult `json:"job_submit_result,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o CreateExecuteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecuteJobResponse struct{}"
	}

	return strings.Join([]string{"CreateExecuteJobResponse", string(data)}, " ")
}

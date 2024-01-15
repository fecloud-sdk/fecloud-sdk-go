package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateReassignmentTaskResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReassignmentTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReassignmentTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateReassignmentTaskResponse", string(data)}, " ")
}

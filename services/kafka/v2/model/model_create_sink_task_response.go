package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSinkTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSinkTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSinkTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSinkTaskResponse", string(data)}, " ")
}

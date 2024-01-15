package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSinkTaskRequest struct {
	ConnectorId string `json:"connector_id"`

	Body *CreateSinkTaskReq `json:"body,omitempty"`
}

func (o CreateSinkTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSinkTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSinkTaskRequest", string(data)}, " ")
}

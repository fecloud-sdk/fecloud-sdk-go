package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSinkTaskRequest struct {
	ConnectorId string `json:"connector_id"`

	TaskId string `json:"task_id"`
}

func (o DeleteSinkTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSinkTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteSinkTaskRequest", string(data)}, " ")
}

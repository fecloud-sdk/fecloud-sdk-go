package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSinkTasksRequest struct {
	ConnectorId string `json:"connector_id"`
}

func (o ListSinkTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSinkTasksRequest struct{}"
	}

	return strings.Join([]string{"ListSinkTasksRequest", string(data)}, " ")
}

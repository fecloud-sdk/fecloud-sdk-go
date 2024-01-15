package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteConnectorRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DeleteConnectorRequestBody `json:"body,omitempty"`
}

func (o DeleteConnectorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectorRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnectorRequest", string(data)}, " ")
}

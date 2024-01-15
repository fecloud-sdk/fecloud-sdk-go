package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDeleteConnectorOrderRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ConnectorOrderRequestBody `json:"body,omitempty"`
}

func (o CreateDeleteConnectorOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeleteConnectorOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateDeleteConnectorOrderRequest", string(data)}, " ")
}

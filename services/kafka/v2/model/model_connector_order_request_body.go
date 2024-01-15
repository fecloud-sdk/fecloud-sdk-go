package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConnectorOrderRequestBody struct {
	InstanceId string `json:"instance_id"`

	Url *string `json:"url,omitempty"`
}

func (o ConnectorOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectorOrderRequestBody struct{}"
	}

	return strings.Join([]string{"ConnectorOrderRequestBody", string(data)}, " ")
}

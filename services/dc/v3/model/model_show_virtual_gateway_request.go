package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowVirtualGatewayRequest struct {
	Fields *[]string `json:"fields,omitempty"`

	VirtualGatewayId string `json:"virtual_gateway_id"`
}

func (o ShowVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowVirtualGatewayRequest", string(data)}, " ")
}

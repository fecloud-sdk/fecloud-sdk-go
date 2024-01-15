package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVirtualGatewayRequest struct {
	VirtualGatewayId string `json:"virtual_gateway_id"`
}

func (o DeleteVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"DeleteVirtualGatewayRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVirtualGatewayRequest struct {
	VirtualGatewayId string `json:"virtual_gateway_id"`

	Body *UpdateVirtualGatewayRequestBody `json:"body,omitempty"`
}

func (o UpdateVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGatewayRequest", string(data)}, " ")
}

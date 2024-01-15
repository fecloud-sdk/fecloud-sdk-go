package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualGatewayRequestBody struct {
	VirtualGateway *CreateVirtualGateway `json:"virtual_gateway,omitempty"`
}

func (o CreateVirtualGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayRequestBody", string(data)}, " ")
}

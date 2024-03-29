package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualGatewayRequest struct {
	Body *CreateVirtualGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayRequest", string(data)}, " ")
}

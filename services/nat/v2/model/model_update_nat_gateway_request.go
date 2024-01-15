package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewayRequest struct {
	NatGatewayId string `json:"nat_gateway_id"`

	Body *UpdateNatGatewayRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteNatGatewayRequest struct {
	NatGatewayId string `json:"nat_gateway_id"`
}

func (o DeleteNatGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayRequest", string(data)}, " ")
}

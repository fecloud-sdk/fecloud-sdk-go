package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewayResponse struct {
	NatGateway     *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayResponse", string(data)}, " ")
}

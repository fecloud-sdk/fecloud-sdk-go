package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteNatGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayResponse", string(data)}, " ")
}

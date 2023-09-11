package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteNatGatewayDnatRuleResponse Response Object
type DeleteNatGatewayDnatRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewayDnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayDnatRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayDnatRuleResponse", string(data)}, " ")
}

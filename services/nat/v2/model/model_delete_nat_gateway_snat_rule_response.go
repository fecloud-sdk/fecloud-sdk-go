package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteNatGatewaySnatRuleResponse Response Object
type DeleteNatGatewaySnatRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewaySnatRuleResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteNatGatewaySnatRuleRequest struct {
	NatGatewayId string `json:"nat_gateway_id"`

	SnatRuleId string `json:"snat_rule_id"`
}

func (o DeleteNatGatewaySnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewaySnatRuleRequest", string(data)}, " ")
}

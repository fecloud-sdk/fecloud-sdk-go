package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowNatGatewaySnatRuleResponse struct {
	SnatRule       *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowNatGatewaySnatRuleResponse", string(data)}, " ")
}

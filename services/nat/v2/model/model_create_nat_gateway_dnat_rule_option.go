package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewayDnatRuleOption struct {
	DnatRule *CreateNatGatewayDnatOption `json:"dnat_rule"`
}

func (o CreateNatGatewayDnatRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayDnatRuleOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayDnatRuleOption", string(data)}, " ")
}

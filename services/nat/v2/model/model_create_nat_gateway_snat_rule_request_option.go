package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewaySnatRuleRequestOption struct {
	SnatRule *CreateNatGatewaySnatRuleOption `json:"snat_rule"`
}

func (o CreateNatGatewaySnatRuleRequestOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleRequestOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleRequestOption", string(data)}, " ")
}

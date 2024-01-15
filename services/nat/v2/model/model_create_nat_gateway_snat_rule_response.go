package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewaySnatRuleResponse struct {
	SnatRule       *CreateNatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o CreateNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleResponse", string(data)}, " ")
}

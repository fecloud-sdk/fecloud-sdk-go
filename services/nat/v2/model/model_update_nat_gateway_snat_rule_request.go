package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewaySnatRuleRequest struct {
	SnatRuleId string `json:"snat_rule_id"`

	Body *UpdateNatGatewaySnatRuleRequestOption `json:"body,omitempty"`
}

func (o UpdateNatGatewaySnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleRequest", string(data)}, " ")
}

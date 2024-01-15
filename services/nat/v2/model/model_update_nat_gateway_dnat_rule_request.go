package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateNatGatewayDnatRuleRequest struct {
	DnatRuleId string `json:"dnat_rule_id"`

	Body *UpdateNatGatewayDnatRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequest", string(data)}, " ")
}

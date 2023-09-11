package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateNatGatewayDnatRuleRequestBody 更新DNAT规则的响应体。
type UpdateNatGatewayDnatRuleRequestBody struct {
	DnatRule *UpdateNatGatewayDnatRuleOption `json:"dnat_rule"`
}

func (o UpdateNatGatewayDnatRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequestBody", string(data)}, " ")
}

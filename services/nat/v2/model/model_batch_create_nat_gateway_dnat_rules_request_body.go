package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateNatGatewayDnatRulesRequestBody struct {
	DnatRules []CreateNatGatewayDnatOption `json:"dnat_rules"`
}

func (o BatchCreateNatGatewayDnatRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNatGatewayDnatRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesRequestBody", string(data)}, " ")
}

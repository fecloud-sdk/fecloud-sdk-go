package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSecurityGroupRuleRequest struct {
	SecurityGroupRuleId string `json:"security_group_rule_id"`
}

func (o DeleteSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupRuleRequest", string(data)}, " ")
}

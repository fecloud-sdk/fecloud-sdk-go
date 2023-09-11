package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRuleResponse struct {
	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o CreateSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleResponse", string(data)}, " ")
}

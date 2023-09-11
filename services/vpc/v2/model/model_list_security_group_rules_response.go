package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityGroupRulesResponse struct {
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesResponse", string(data)}, " ")
}

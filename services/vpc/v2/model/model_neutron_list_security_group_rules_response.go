package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronListSecurityGroupRulesResponse struct {
	SecurityGroupRules *[]NeutronSecurityGroupRule `json:"security_group_rules,omitempty"`
	HttpStatusCode     int                         `json:"-"`
}

func (o NeutronListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupRulesResponse", string(data)}, " ")
}

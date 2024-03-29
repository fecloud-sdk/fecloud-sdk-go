package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallRuleRequestBody struct {
	FirewallRule *NeutronCreateFirewallRuleOption `json:"firewall_rule"`
}

func (o NeutronCreateFirewallRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleRequestBody", string(data)}, " ")
}

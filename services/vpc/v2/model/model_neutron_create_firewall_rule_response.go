package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallRuleResponse struct {
	FirewallRule   *NeutronFirewallRule `json:"firewall_rule,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o NeutronCreateFirewallRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleResponse", string(data)}, " ")
}

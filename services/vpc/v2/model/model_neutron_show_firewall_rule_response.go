package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronShowFirewallRuleResponse struct {
	FirewallRule   *NeutronFirewallRule `json:"firewall_rule,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o NeutronShowFirewallRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallRuleResponse", string(data)}, " ")
}

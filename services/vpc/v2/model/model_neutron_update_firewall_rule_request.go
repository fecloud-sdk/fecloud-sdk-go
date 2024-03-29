package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallRuleRequest struct {
	FirewallRuleId string `json:"firewall_rule_id"`

	Body *NeutronUpdateFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallRuleRequest", string(data)}, " ")
}

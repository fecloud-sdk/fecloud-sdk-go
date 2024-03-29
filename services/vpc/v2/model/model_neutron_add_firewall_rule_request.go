package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronAddFirewallRuleRequest struct {
	FirewallPolicyId string `json:"firewall_policy_id"`

	Body *NeutronInsertFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronAddFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronAddFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronAddFirewallRuleRequest", string(data)}, " ")
}

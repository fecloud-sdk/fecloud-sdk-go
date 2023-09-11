package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallPolicyRequestBody struct {
	FirewallPolicy *NeutronCreateFirewallPolicyOption `json:"firewall_policy"`
}

func (o NeutronCreateFirewallPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallPolicyRequestBody", string(data)}, " ")
}

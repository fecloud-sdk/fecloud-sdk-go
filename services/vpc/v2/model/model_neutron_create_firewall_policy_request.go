package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallPolicyRequest struct {
	Body *NeutronCreateFirewallPolicyRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallPolicyRequest", string(data)}, " ")
}

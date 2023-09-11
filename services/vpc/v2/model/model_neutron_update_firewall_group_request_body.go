package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronUpdateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronUpdateFirewallGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupRequestBody", string(data)}, " ")
}

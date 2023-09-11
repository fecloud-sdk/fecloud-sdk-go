package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallGroupResponse struct {
	FirewallGroup  *NeutronFirewallGroup `json:"firewall_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronUpdateFirewallGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupResponse", string(data)}, " ")
}

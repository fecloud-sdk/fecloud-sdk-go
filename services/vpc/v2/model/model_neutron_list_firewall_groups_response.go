package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronListFirewallGroupsResponse struct {
	FirewallGroups *[]NeutronFirewallGroup `json:"firewall_groups,omitempty"`

	FirewallGroupsLinks *[]NeutronPageLink `json:"firewall_groups_links,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o NeutronListFirewallGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallGroupsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallGroupsResponse", string(data)}, " ")
}

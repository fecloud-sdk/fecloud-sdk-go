package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VirtualGateway struct {
	Id *string `json:"id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *string `json:"type,omitempty"`

	LocalEpGroup *[]string `json:"local_ep_group,omitempty"`

	LocalEpGroupIpv6 *[]string `json:"local_ep_group_ipv6,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Status *string `json:"status,omitempty"`

	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o VirtualGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualGateway struct{}"
	}

	return strings.Join([]string{"VirtualGateway", string(data)}, " ")
}

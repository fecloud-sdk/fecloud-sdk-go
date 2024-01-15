package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualGateway struct {
	VpcId string `json:"vpc_id"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	LocalEpGroup []string `json:"local_ep_group"`

	LocalEpGroupIpv6 *[]string `json:"local_ep_group_ipv6,omitempty"`

	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateVirtualGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGateway struct{}"
	}

	return strings.Join([]string{"CreateVirtualGateway", string(data)}, " ")
}

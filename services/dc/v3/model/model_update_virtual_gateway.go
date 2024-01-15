package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateVirtualGateway struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	LocalEpGroup *[]string `json:"local_ep_group,omitempty"`

	LocalEpGroupIpv6 *[]string `json:"local_ep_group_ipv6,omitempty"`
}

func (o UpdateVirtualGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGateway struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGateway", string(data)}, " ")
}

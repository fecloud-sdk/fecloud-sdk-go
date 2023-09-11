package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVpcPeeringsResponse struct {
	Peerings *[]VpcPeering `json:"peerings,omitempty"`

	PeeringsLinks  *[]NeutronPageLink `json:"peerings_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVpcPeeringsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcPeeringsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcPeeringsResponse", string(data)}, " ")
}

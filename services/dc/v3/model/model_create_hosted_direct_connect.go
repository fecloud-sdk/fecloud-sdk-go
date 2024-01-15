package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHostedDirectConnect struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Bandwidth int32 `json:"bandwidth"`

	HostingId string `json:"hosting_id"`

	Vlan int32 `json:"vlan"`

	ResourceTenantId string `json:"resource_tenant_id"`

	PeerLocation *string `json:"peer_location,omitempty"`
}

func (o CreateHostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnect struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnect", string(data)}, " ")
}

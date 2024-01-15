package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPublicZonesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Zones *[]PublicZoneResp `json:"zones,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPublicZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPublicZonesResponse", string(data)}, " ")
}

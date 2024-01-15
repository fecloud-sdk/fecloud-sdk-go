package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateZonesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Zones          *[]PrivateZoneResp `json:"zones,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListPrivateZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesResponse", string(data)}, " ")
}

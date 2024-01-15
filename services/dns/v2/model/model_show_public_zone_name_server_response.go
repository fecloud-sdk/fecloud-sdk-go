package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPublicZoneNameServerResponse struct {
	Nameservers    *[]Nameserver `json:"nameservers,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowPublicZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerResponse", string(data)}, " ")
}

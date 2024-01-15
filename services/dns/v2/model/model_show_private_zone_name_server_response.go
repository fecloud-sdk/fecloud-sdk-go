package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateZoneNameServerResponse struct {
	Nameservers    *[]PrivateNameServer `json:"nameservers,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPrivateZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerResponse", string(data)}, " ")
}

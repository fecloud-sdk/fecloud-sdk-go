package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociatedTransitIp struct {
	TransitIpId *string `json:"transit_ip_id,omitempty"`

	TransitIpAddress *string `json:"transit_ip_address,omitempty"`
}

func (o AssociatedTransitIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedTransitIp struct{}"
	}

	return strings.Join([]string{"AssociatedTransitIp", string(data)}, " ")
}

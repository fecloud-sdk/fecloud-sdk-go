package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAvailableZonesRespAvailableZones struct {
	SoldOut *bool `json:"soldOut,omitempty"`

	Id *string `json:"id,omitempty"`

	Code *string `json:"code,omitempty"`

	Name *string `json:"name,omitempty"`

	Port *string `json:"port,omitempty"`

	ResourceAvailability *string `json:"resource_availability,omitempty"`

	DefaultAz *bool `json:"default_az,omitempty"`

	RemainTime *int64 `json:"remain_time,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o ListAvailableZonesRespAvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesRespAvailableZones struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesRespAvailableZones", string(data)}, " ")
}

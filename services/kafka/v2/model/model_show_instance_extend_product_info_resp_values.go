package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceExtendProductInfoRespValues struct {
	Detail *[]ShowInstanceExtendProductInfoRespDetail `json:"detail,omitempty"`

	Name *string `json:"name,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`
}

func (o ShowInstanceExtendProductInfoRespValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoRespValues struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoRespValues", string(data)}, " ")
}

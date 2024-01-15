package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProductsRespValues struct {
	Detail *[]ListProductsRespDetail `json:"detail,omitempty"`

	Name *string `json:"name,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`
}

func (o ListProductsRespValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespValues struct{}"
	}

	return strings.Join([]string{"ListProductsRespValues", string(data)}, " ")
}

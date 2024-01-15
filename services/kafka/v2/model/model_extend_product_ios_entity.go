package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExtendProductIosEntity struct {
	IoSpec *string `json:"io_spec,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	Type *string `json:"type,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`
}

func (o ExtendProductIosEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendProductIosEntity struct{}"
	}

	return strings.Join([]string{"ExtendProductIosEntity", string(data)}, " ")
}

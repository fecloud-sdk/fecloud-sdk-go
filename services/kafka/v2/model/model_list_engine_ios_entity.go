package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEngineIosEntity struct {
	IoSpec *string `json:"io_spec,omitempty"`

	Type *string `json:"type,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`
}

func (o ListEngineIosEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineIosEntity struct{}"
	}

	return strings.Join([]string{"ListEngineIosEntity", string(data)}, " ")
}

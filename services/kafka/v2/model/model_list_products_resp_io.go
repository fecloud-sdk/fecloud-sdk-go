package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProductsRespIo struct {
	IoType *string `json:"io_type,omitempty"`

	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	VolumeType *string `json:"volume_type,omitempty"`
}

func (o ListProductsRespIo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespIo struct{}"
	}

	return strings.Join([]string{"ListProductsRespIo", string(data)}, " ")
}

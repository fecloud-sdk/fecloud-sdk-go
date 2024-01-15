package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProductsRespDetail struct {
	Tps *string `json:"tps,omitempty"`

	Storage *string `json:"storage,omitempty"`

	PartitionNum *string `json:"partition_num,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	Io *[]ListProductsRespIo `json:"io,omitempty"`

	Bandwidth *string `json:"bandwidth,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	ArchType *string `json:"arch_type,omitempty"`
}

func (o ListProductsRespDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespDetail struct{}"
	}

	return strings.Join([]string{"ListProductsRespDetail", string(data)}, " ")
}

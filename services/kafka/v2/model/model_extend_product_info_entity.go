package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExtendProductInfoEntity struct {
	Type *string `json:"type,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	ArchTypes *[]string `json:"arch_types,omitempty"`

	ChargingMode *[]string `json:"charging_mode,omitempty"`

	Ios *[]ExtendProductIosEntity `json:"ios,omitempty"`

	SupportFeatures *[]ExtendProductSupportFeaturesEntity `json:"support_features,omitempty"`

	Properties *ExtendProductPropertiesEntity `json:"properties,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`
}

func (o ExtendProductInfoEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendProductInfoEntity struct{}"
	}

	return strings.Join([]string{"ExtendProductInfoEntity", string(data)}, " ")
}

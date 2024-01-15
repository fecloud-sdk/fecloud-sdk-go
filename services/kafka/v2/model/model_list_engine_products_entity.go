package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEngineProductsEntity struct {
	Type *string `json:"type,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	BillingCode *string `json:"billing_code,omitempty"`

	ArchTypes *[]string `json:"arch_types,omitempty"`

	ChargingMode *[]string `json:"charging_mode,omitempty"`

	Ios *[]ListEngineIosEntity `json:"ios,omitempty"`

	SupportFeatures *[]ListEngineSupportFeaturesEntity `json:"support_features,omitempty"`

	Properties *ListEnginePropertiesEntity `json:"properties,omitempty"`
}

func (o ListEngineProductsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsEntity struct{}"
	}

	return strings.Join([]string{"ListEngineProductsEntity", string(data)}, " ")
}

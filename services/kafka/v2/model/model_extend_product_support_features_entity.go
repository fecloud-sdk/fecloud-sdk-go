package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExtendProductSupportFeaturesEntity struct {
	Name *string `json:"name,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
}

func (o ExtendProductSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendProductSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"ExtendProductSupportFeaturesEntity", string(data)}, " ")
}

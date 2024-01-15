package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEngineSupportFeaturesEntity struct {
	Name *string `json:"name,omitempty"`

	Properties *ListEngineSupportFeaturesPropertiesEntity `json:"properties,omitempty"`
}

func (o ListEngineSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"ListEngineSupportFeaturesEntity", string(data)}, " ")
}

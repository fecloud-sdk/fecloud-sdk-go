package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEngineSupportFeaturesPropertiesEntity struct {
	MaxTask *string `json:"max_task,omitempty"`

	MinTask *string `json:"min_task,omitempty"`

	MaxNode *string `json:"max_node,omitempty"`

	MinNode *string `json:"min_node,omitempty"`
}

func (o ListEngineSupportFeaturesPropertiesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineSupportFeaturesPropertiesEntity struct{}"
	}

	return strings.Join([]string{"ListEngineSupportFeaturesPropertiesEntity", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TopicEntityTopicOtherConfigs struct {
	Name *string `json:"name,omitempty"`

	ValidValues *string `json:"valid_values,omitempty"`

	DefaultValue *string `json:"default_value,omitempty"`

	ConfigType *string `json:"config_type,omitempty"`

	Value *string `json:"value,omitempty"`

	ValueType *string `json:"value_type,omitempty"`
}

func (o TopicEntityTopicOtherConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicEntityTopicOtherConfigs struct{}"
	}

	return strings.Join([]string{"TopicEntityTopicOtherConfigs", string(data)}, " ")
}

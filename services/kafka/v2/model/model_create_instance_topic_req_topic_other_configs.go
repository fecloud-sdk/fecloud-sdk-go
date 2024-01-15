package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceTopicReqTopicOtherConfigs struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o CreateInstanceTopicReqTopicOtherConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicReqTopicOtherConfigs struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicReqTopicOtherConfigs", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DiskusageTopicEntity struct {
	Size *string `json:"size,omitempty"`

	TopicName *string `json:"topic_name,omitempty"`

	TopicPartition *string `json:"topic_partition,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`
}

func (o DiskusageTopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskusageTopicEntity struct{}"
	}

	return strings.Join([]string{"DiskusageTopicEntity", string(data)}, " ")
}

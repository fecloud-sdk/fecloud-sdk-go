package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPartitionMessageEntity struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`

	Topic *string `json:"topic,omitempty"`

	Partition *int32 `json:"partition,omitempty"`

	MessageOffset *int64 `json:"message_offset,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o ShowPartitionMessageEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionMessageEntity struct{}"
	}

	return strings.Join([]string{"ShowPartitionMessageEntity", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type KafkaTopicPartitionResponsePartitions struct {
	Partition *int32 `json:"partition,omitempty"`

	StartOffset *int64 `json:"start_offset,omitempty"`

	LastOffset *int64 `json:"last_offset,omitempty"`

	MessageCount *int64 `json:"message_count,omitempty"`

	LastUpdateTime *int64 `json:"last_update_time,omitempty"`
}

func (o KafkaTopicPartitionResponsePartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicPartitionResponsePartitions struct{}"
	}

	return strings.Join([]string{"KafkaTopicPartitionResponsePartitions", string(data)}, " ")
}

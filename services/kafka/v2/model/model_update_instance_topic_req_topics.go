package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceTopicReqTopics struct {
	Id string `json:"id"`

	RetentionTime *int32 `json:"retention_time,omitempty"`

	SyncReplication *bool `json:"sync_replication,omitempty"`

	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	NewPartitionNumbers *int32 `json:"new_partition_numbers,omitempty"`

	TopicOtherConfigs *[]CreateInstanceTopicReqTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	TopicDesc *string `json:"topic_desc,omitempty"`
}

func (o UpdateInstanceTopicReqTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReqTopics struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReqTopics", string(data)}, " ")
}

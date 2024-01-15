package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TopicEntity struct {
	PoliciesOnly *bool `json:"policiesOnly,omitempty"`

	Name *string `json:"name,omitempty"`

	Replication *int32 `json:"replication,omitempty"`

	Partition *int32 `json:"partition,omitempty"`

	RetentionTime *int32 `json:"retention_time,omitempty"`

	SyncReplication *bool `json:"sync_replication,omitempty"`

	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	ExternalConfigs *interface{} `json:"external_configs,omitempty"`

	TopicType *int32 `json:"topic_type,omitempty"`

	TopicOtherConfigs *[]TopicEntityTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	TopicDesc *string `json:"topic_desc,omitempty"`

	CreatedAt *int64 `json:"created_at,omitempty"`
}

func (o TopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicEntity struct{}"
	}

	return strings.Join([]string{"TopicEntity", string(data)}, " ")
}

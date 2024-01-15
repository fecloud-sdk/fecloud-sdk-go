package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceTopicReq struct {
	Id string `json:"id"`

	Replication *int32 `json:"replication,omitempty"`

	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	Partition *int32 `json:"partition,omitempty"`

	SyncReplication *bool `json:"sync_replication,omitempty"`

	RetentionTime *int32 `json:"retention_time,omitempty"`

	TopicOtherConfigs *[]CreateInstanceTopicReqTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	TopicDesc *string `json:"topic_desc,omitempty"`
}

func (o CreateInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicReq", string(data)}, " ")
}

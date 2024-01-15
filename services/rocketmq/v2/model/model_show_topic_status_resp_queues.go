package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTopicStatusRespQueues struct {
	Id *int32 `json:"id,omitempty"`

	MinOffset *int32 `json:"min_offset,omitempty"`

	MaxOffset *int32 `json:"max_offset,omitempty"`

	LastMessageTime *int64 `json:"last_message_time,omitempty"`
}

func (o ShowTopicStatusRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRespQueues struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRespQueues", string(data)}, " ")
}

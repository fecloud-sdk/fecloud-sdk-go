package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceTopicsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Size *int32 `json:"size,omitempty"`

	RemainPartitions *int32 `json:"remain_partitions,omitempty"`

	MaxPartitions *int32 `json:"max_partitions,omitempty"`

	TopicMaxPartitions *int32 `json:"topic_max_partitions,omitempty"`

	Topics         *[]TopicEntity `json:"topics,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTopicPartitionsRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTopicPartitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicPartitionsRequest struct{}"
	}

	return strings.Join([]string{"ListTopicPartitionsRequest", string(data)}, " ")
}

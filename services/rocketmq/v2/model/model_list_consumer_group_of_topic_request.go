package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConsumerGroupOfTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListConsumerGroupOfTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupOfTopicRequest struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupOfTopicRequest", string(data)}, " ")
}

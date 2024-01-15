package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTopicProducersRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTopicProducersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicProducersRequest struct{}"
	}

	return strings.Join([]string{"ListTopicProducersRequest", string(data)}, " ")
}

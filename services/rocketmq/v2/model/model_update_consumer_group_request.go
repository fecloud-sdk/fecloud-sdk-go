package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConsumerGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`

	Body *UpdateConsumerGroup `json:"body,omitempty"`
}

func (o UpdateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateConsumerGroupRequest", string(data)}, " ")
}

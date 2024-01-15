package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateConsumerGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchUpdateConsumerGroupReq `json:"body,omitempty"`
}

func (o BatchUpdateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupRequest", string(data)}, " ")
}

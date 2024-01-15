package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateKafkaConsumerGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateGroupReq `json:"body,omitempty"`
}

func (o CreateKafkaConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateKafkaConsumerGroupRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConsumerGroupOrBatchDeleteConsumerGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Action *string `json:"action,omitempty"`

	Body *CreateConsumerGroupOrBatchDeleteConsumerGroupReq `json:"body,omitempty"`
}

func (o CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupOrBatchDeleteConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupOrBatchDeleteConsumerGroupRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPartitionBeginningMessageRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Partition int32 `json:"partition"`
}

func (o ShowPartitionBeginningMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionBeginningMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionBeginningMessageRequest", string(data)}, " ")
}

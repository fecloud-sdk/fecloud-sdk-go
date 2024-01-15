package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPartitionEndMessageRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Partition int32 `json:"partition"`
}

func (o ShowPartitionEndMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionEndMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionEndMessageRequest", string(data)}, " ")
}

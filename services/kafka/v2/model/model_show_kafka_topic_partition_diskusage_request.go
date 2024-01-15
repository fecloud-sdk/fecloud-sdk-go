package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKafkaTopicPartitionDiskusageRequest struct {
	InstanceId string `json:"instance_id"`

	MinSize *string `json:"minSize,omitempty"`

	Top *string `json:"top,omitempty"`

	Percentage *string `json:"percentage,omitempty"`
}

func (o ShowKafkaTopicPartitionDiskusageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicPartitionDiskusageRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicPartitionDiskusageRequest", string(data)}, " ")
}

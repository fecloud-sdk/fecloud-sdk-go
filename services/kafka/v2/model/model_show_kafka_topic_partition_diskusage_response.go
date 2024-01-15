package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKafkaTopicPartitionDiskusageResponse struct {
	BrokerList     *[]DiskusageEntity `json:"broker_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowKafkaTopicPartitionDiskusageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicPartitionDiskusageResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicPartitionDiskusageResponse", string(data)}, " ")
}

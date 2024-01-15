package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTopicPartitionsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Partitions     *[]KafkaTopicPartitionResponsePartitions `json:"partitions,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListTopicPartitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicPartitionsResponse struct{}"
	}

	return strings.Join([]string{"ListTopicPartitionsResponse", string(data)}, " ")
}

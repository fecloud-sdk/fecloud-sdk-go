package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type KafkaTopicProducerResponseProducers struct {
	ProducerAddress *string `json:"producer_address,omitempty"`

	BrokerAddress *string `json:"broker_address,omitempty"`

	JoinTime *int64 `json:"join_time,omitempty"`
}

func (o KafkaTopicProducerResponseProducers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicProducerResponseProducers struct{}"
	}

	return strings.Join([]string{"KafkaTopicProducerResponseProducers", string(data)}, " ")
}

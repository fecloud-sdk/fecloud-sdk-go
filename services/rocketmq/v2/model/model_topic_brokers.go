package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TopicBrokers struct {
	BrokerName *string `json:"broker_name,omitempty"`

	ReadQueueNum float32 `json:"read_queue_num,omitempty"`

	WriteQueueNum float32 `json:"write_queue_num,omitempty"`
}

func (o TopicBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicBrokers struct{}"
	}

	return strings.Join([]string{"TopicBrokers", string(data)}, " ")
}

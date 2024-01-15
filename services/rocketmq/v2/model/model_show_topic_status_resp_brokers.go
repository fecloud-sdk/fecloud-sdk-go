package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTopicStatusRespBrokers struct {
	Queues *[]ShowTopicStatusRespQueues `json:"queues,omitempty"`

	BrokerName *string `json:"broker_name,omitempty"`
}

func (o ShowTopicStatusRespBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRespBrokers struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRespBrokers", string(data)}, " ")
}

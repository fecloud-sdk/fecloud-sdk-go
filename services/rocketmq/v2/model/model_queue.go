package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Queue struct {
	Id *int32 `json:"id,omitempty"`

	Lag *int64 `json:"lag,omitempty"`

	BrokerOffset *int64 `json:"broker_offset,omitempty"`

	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	LastMessageTime *int64 `json:"last_message_time,omitempty"`
}

func (o Queue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Queue struct{}"
	}

	return strings.Join([]string{"Queue", string(data)}, " ")
}

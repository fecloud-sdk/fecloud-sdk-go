package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConsumerGroup struct {
	Enabled *bool `json:"enabled,omitempty"`

	Broadcast *bool `json:"broadcast,omitempty"`

	Brokers *[]string `json:"brokers,omitempty"`

	Name *string `json:"name,omitempty"`

	GroupDesc *string `json:"group_desc,omitempty"`

	RetryMaxTime float32 `json:"retry_max_time,omitempty"`

	FromBeginning *bool `json:"from_beginning,omitempty"`
}

func (o ConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroup struct{}"
	}

	return strings.Join([]string{"ConsumerGroup", string(data)}, " ")
}

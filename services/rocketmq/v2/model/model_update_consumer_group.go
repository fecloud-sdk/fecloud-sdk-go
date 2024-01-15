package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConsumerGroup struct {
	Enabled bool `json:"enabled"`

	Broadcast bool `json:"broadcast"`

	Brokers *[]string `json:"brokers,omitempty"`

	Name *string `json:"name,omitempty"`

	RetryMaxTime float32 `json:"retry_max_time"`

	FromBeginning *bool `json:"from_beginning,omitempty"`
}

func (o UpdateConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsumerGroup struct{}"
	}

	return strings.Join([]string{"UpdateConsumerGroup", string(data)}, " ")
}

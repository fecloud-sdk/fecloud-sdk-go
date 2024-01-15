package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Brokers struct {
	BrokerName *string `json:"broker_name,omitempty"`

	Queues *[]Queue `json:"queues,omitempty"`
}

func (o Brokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Brokers struct{}"
	}

	return strings.Join([]string{"Brokers", string(data)}, " ")
}

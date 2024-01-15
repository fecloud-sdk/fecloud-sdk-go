package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBrokersRespBrokers struct {
	Ids *[]float32 `json:"ids,omitempty"`

	BrokerName *string `json:"broker_name,omitempty"`
}

func (o ListBrokersRespBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersRespBrokers struct{}"
	}

	return strings.Join([]string{"ListBrokersRespBrokers", string(data)}, " ")
}

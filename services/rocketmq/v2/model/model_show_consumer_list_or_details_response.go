package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConsumerListOrDetailsResponse struct {
	Topics *[]string `json:"topics,omitempty"`

	Total *int32 `json:"total,omitempty"`

	Lag *int64 `json:"lag,omitempty"`

	MaxOffset *int64 `json:"max_offset,omitempty"`

	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	Brokers        *[]Brokers `json:"brokers,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowConsumerListOrDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsResponse", string(data)}, " ")
}

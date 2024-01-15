package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConsumerDetailResp struct {
	Lag *int64 `json:"lag,omitempty"`

	MaxOffset *int64 `json:"max_offset,omitempty"`

	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	Brokers *[]Brokers `json:"brokers,omitempty"`
}

func (o ConsumerDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerDetailResp struct{}"
	}

	return strings.Join([]string{"ConsumerDetailResp", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConsumerListOrDetailsRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`

	Topic *string `json:"topic,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowConsumerListOrDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsRequest", string(data)}, " ")
}

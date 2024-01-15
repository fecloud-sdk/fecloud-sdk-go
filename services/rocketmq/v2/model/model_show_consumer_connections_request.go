package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConsumerConnectionsRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	IsDetail *bool `json:"is_detail,omitempty"`
}

func (o ShowConsumerConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerConnectionsRequest", string(data)}, " ")
}

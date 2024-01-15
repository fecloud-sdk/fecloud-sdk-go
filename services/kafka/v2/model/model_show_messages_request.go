package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMessagesRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Partition *string `json:"partition,omitempty"`
}

func (o ShowMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessagesRequest struct{}"
	}

	return strings.Join([]string{"ShowMessagesRequest", string(data)}, " ")
}

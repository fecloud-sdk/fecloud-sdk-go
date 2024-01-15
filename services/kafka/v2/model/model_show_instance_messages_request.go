package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceMessagesRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Asc *bool `json:"asc,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Download *bool `json:"download,omitempty"`

	MessageOffset *string `json:"message_offset,omitempty"`

	Partition *string `json:"partition,omitempty"`

	Keyword *string `json:"keyword,omitempty"`
}

func (o ShowInstanceMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMessagesRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceMessagesRequest", string(data)}, " ")
}

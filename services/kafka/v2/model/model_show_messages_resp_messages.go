package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMessagesRespMessages struct {
	Topic *string `json:"topic,omitempty"`

	Partition *int32 `json:"partition,omitempty"`

	MessageOffset *int32 `json:"message_offset,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o ShowMessagesRespMessages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessagesRespMessages struct{}"
	}

	return strings.Join([]string{"ShowMessagesRespMessages", string(data)}, " ")
}

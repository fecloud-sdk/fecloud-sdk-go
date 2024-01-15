package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MessagesEntity struct {
	Topic *string `json:"topic,omitempty"`

	Partition *int32 `json:"partition,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`

	HugeMessage *bool `json:"huge_message,omitempty"`

	MessageOffset *int32 `json:"message_offset,omitempty"`

	MessageId *string `json:"message_id,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (o MessagesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessagesEntity struct{}"
	}

	return strings.Join([]string{"MessagesEntity", string(data)}, " ")
}

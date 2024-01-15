package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Message struct {
	MsgId *string `json:"msg_id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Topic *string `json:"topic,omitempty"`

	StoreTimestamp float32 `json:"store_timestamp,omitempty"`

	BornTimestamp float32 `json:"born_timestamp,omitempty"`

	ReconsumeTimes *int32 `json:"reconsume_times,omitempty"`

	Body *string `json:"body,omitempty"`

	BodyCrc float32 `json:"body_crc,omitempty"`

	StoreSize float32 `json:"store_size,omitempty"`

	PropertyList *[]MessagePropertyList `json:"property_list,omitempty"`

	BornHost *string `json:"born_host,omitempty"`

	StoreHost *string `json:"store_host,omitempty"`

	QueueId *int32 `json:"queue_id,omitempty"`

	QueueOffset *int32 `json:"queue_offset,omitempty"`
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}

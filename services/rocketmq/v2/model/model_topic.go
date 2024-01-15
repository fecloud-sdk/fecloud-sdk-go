package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type Topic struct {
	Name *string `json:"name,omitempty"`

	TotalReadQueueNum float32 `json:"total_read_queue_num,omitempty"`

	TotalWriteQueueNum float32 `json:"total_write_queue_num,omitempty"`

	Permission *TopicPermission `json:"permission,omitempty"`

	Brokers *[]TopicBrokers `json:"brokers,omitempty"`

	MessageType *TopicMessageType `json:"message_type,omitempty"`
}

func (o Topic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topic struct{}"
	}

	return strings.Join([]string{"Topic", string(data)}, " ")
}

type TopicPermission struct {
	value string
}

type TopicPermissionEnum struct {
	SUB TopicPermission
	PUB TopicPermission
	ALL TopicPermission
}

func GetTopicPermissionEnum() TopicPermissionEnum {
	return TopicPermissionEnum{
		SUB: TopicPermission{
			value: "sub",
		},
		PUB: TopicPermission{
			value: "pub",
		},
		ALL: TopicPermission{
			value: "all",
		},
	}
}

func (c TopicPermission) Value() string {
	return c.value
}

func (c TopicPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopicPermission) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type TopicMessageType struct {
	value string
}

type TopicMessageTypeEnum struct {
	NORMAL      TopicMessageType
	FIFO        TopicMessageType
	DELAY       TopicMessageType
	TRANSACTION TopicMessageType
}

func GetTopicMessageTypeEnum() TopicMessageTypeEnum {
	return TopicMessageTypeEnum{
		NORMAL: TopicMessageType{
			value: "NORMAL",
		},
		FIFO: TopicMessageType{
			value: "FIFO",
		},
		DELAY: TopicMessageType{
			value: "DELAY",
		},
		TRANSACTION: TopicMessageType{
			value: "TRANSACTION",
		},
	}
}

func (c TopicMessageType) Value() string {
	return c.value
}

func (c TopicMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopicMessageType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

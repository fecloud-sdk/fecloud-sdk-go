package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListMessagesRequest struct {
	Engine ListMessagesRequestEngine `json:"engine"`

	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Key *string `json:"key,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	MsgId *string `json:"msg_id,omitempty"`
}

func (o ListMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesRequest struct{}"
	}

	return strings.Join([]string{"ListMessagesRequest", string(data)}, " ")
}

type ListMessagesRequestEngine struct {
	value string
}

type ListMessagesRequestEngineEnum struct {
	RELIABILITY ListMessagesRequestEngine
}

func GetListMessagesRequestEngineEnum() ListMessagesRequestEngineEnum {
	return ListMessagesRequestEngineEnum{
		RELIABILITY: ListMessagesRequestEngine{
			value: "reliability",
		},
	}
}

func (c ListMessagesRequestEngine) Value() string {
	return c.value
}

func (c ListMessagesRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessagesRequestEngine) UnmarshalJSON(b []byte) error {
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

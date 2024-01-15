package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListSlowLogsRequest struct {
	InstanceId string `json:"instance_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	NodeId *string `json:"node_id,omitempty"`

	Type *ListSlowLogsRequestType `json:"type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogsRequest", string(data)}, " ")
}

type ListSlowLogsRequestType struct {
	value string
}

type ListSlowLogsRequestTypeEnum struct {
	INSERT      ListSlowLogsRequestType
	QUERY       ListSlowLogsRequestType
	UPDATE      ListSlowLogsRequestType
	REMOVE      ListSlowLogsRequestType
	GETMORE     ListSlowLogsRequestType
	COMMAND     ListSlowLogsRequestType
	KILLCURSORS ListSlowLogsRequestType
}

func GetListSlowLogsRequestTypeEnum() ListSlowLogsRequestTypeEnum {
	return ListSlowLogsRequestTypeEnum{
		INSERT: ListSlowLogsRequestType{
			value: "INSERT",
		},
		QUERY: ListSlowLogsRequestType{
			value: "QUERY",
		},
		UPDATE: ListSlowLogsRequestType{
			value: "UPDATE",
		},
		REMOVE: ListSlowLogsRequestType{
			value: "REMOVE",
		},
		GETMORE: ListSlowLogsRequestType{
			value: "GETMORE",
		},
		COMMAND: ListSlowLogsRequestType{
			value: "COMMAND",
		},
		KILLCURSORS: ListSlowLogsRequestType{
			value: "KILLCURSORS",
		},
	}
}

func (c ListSlowLogsRequestType) Value() string {
	return c.value
}

func (c ListSlowLogsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowLogsRequestType) UnmarshalJSON(b []byte) error {
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

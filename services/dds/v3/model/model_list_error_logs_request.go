package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListErrorLogsRequest struct {
	InstanceId string `json:"instance_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	NodeId *string `json:"node_id,omitempty"`

	Type *ListErrorLogsRequestType `json:"type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListErrorLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorLogsRequest struct{}"
	}

	return strings.Join([]string{"ListErrorLogsRequest", string(data)}, " ")
}

type ListErrorLogsRequestType struct {
	value string
}

type ListErrorLogsRequestTypeEnum struct {
	WARNING ListErrorLogsRequestType
	ERROR   ListErrorLogsRequestType
}

func GetListErrorLogsRequestTypeEnum() ListErrorLogsRequestTypeEnum {
	return ListErrorLogsRequestTypeEnum{
		WARNING: ListErrorLogsRequestType{
			value: "WARNING",
		},
		ERROR: ListErrorLogsRequestType{
			value: "ERROR",
		},
	}
}

func (c ListErrorLogsRequestType) Value() string {
	return c.value
}

func (c ListErrorLogsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListErrorLogsRequestType) UnmarshalJSON(b []byte) error {
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

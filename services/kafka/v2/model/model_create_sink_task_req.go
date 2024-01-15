package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateSinkTaskReq struct {
	SourceType CreateSinkTaskReqSourceType `json:"source_type"`

	TaskName string `json:"task_name"`

	DestinationType CreateSinkTaskReqDestinationType `json:"destination_type"`

	ObsDestinationDescriptor *ObsDestinationDescriptor `json:"obs_destination_descriptor"`
}

func (o CreateSinkTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSinkTaskReq struct{}"
	}

	return strings.Join([]string{"CreateSinkTaskReq", string(data)}, " ")
}

type CreateSinkTaskReqSourceType struct {
	value string
}

type CreateSinkTaskReqSourceTypeEnum struct {
	BLOB CreateSinkTaskReqSourceType
}

func GetCreateSinkTaskReqSourceTypeEnum() CreateSinkTaskReqSourceTypeEnum {
	return CreateSinkTaskReqSourceTypeEnum{
		BLOB: CreateSinkTaskReqSourceType{
			value: "BLOB",
		},
	}
}

func (c CreateSinkTaskReqSourceType) Value() string {
	return c.value
}

func (c CreateSinkTaskReqSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSinkTaskReqSourceType) UnmarshalJSON(b []byte) error {
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

type CreateSinkTaskReqDestinationType struct {
	value string
}

type CreateSinkTaskReqDestinationTypeEnum struct {
	OBS CreateSinkTaskReqDestinationType
}

func GetCreateSinkTaskReqDestinationTypeEnum() CreateSinkTaskReqDestinationTypeEnum {
	return CreateSinkTaskReqDestinationTypeEnum{
		OBS: CreateSinkTaskReqDestinationType{
			value: "OBS",
		},
	}
}

func (c CreateSinkTaskReqDestinationType) Value() string {
	return c.value
}

func (c CreateSinkTaskReqDestinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSinkTaskReqDestinationType) UnmarshalJSON(b []byte) error {
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

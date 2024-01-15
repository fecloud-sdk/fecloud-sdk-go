package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateRocketMqMigrationTaskRequest struct {
	InstanceId string `json:"instance_id"`

	Overwrite CreateRocketMqMigrationTaskRequestOverwrite `json:"overwrite"`

	Name string `json:"name"`

	Type CreateRocketMqMigrationTaskRequestType `json:"type"`

	Body *string `json:"body,omitempty"`
}

func (o CreateRocketMqMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskRequest", string(data)}, " ")
}

type CreateRocketMqMigrationTaskRequestOverwrite struct {
	value string
}

type CreateRocketMqMigrationTaskRequestOverwriteEnum struct {
	TRUE  CreateRocketMqMigrationTaskRequestOverwrite
	FALSE CreateRocketMqMigrationTaskRequestOverwrite
}

func GetCreateRocketMqMigrationTaskRequestOverwriteEnum() CreateRocketMqMigrationTaskRequestOverwriteEnum {
	return CreateRocketMqMigrationTaskRequestOverwriteEnum{
		TRUE: CreateRocketMqMigrationTaskRequestOverwrite{
			value: "true",
		},
		FALSE: CreateRocketMqMigrationTaskRequestOverwrite{
			value: "false",
		},
	}
}

func (c CreateRocketMqMigrationTaskRequestOverwrite) Value() string {
	return c.value
}

func (c CreateRocketMqMigrationTaskRequestOverwrite) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRocketMqMigrationTaskRequestOverwrite) UnmarshalJSON(b []byte) error {
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

type CreateRocketMqMigrationTaskRequestType struct {
	value string
}

type CreateRocketMqMigrationTaskRequestTypeEnum struct {
	ROCKETMQ         CreateRocketMqMigrationTaskRequestType
	RABBIT_TO_ROCKET CreateRocketMqMigrationTaskRequestType
}

func GetCreateRocketMqMigrationTaskRequestTypeEnum() CreateRocketMqMigrationTaskRequestTypeEnum {
	return CreateRocketMqMigrationTaskRequestTypeEnum{
		ROCKETMQ: CreateRocketMqMigrationTaskRequestType{
			value: "rocketmq",
		},
		RABBIT_TO_ROCKET: CreateRocketMqMigrationTaskRequestType{
			value: "rabbitToRocket",
		},
	}
}

func (c CreateRocketMqMigrationTaskRequestType) Value() string {
	return c.value
}

func (c CreateRocketMqMigrationTaskRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRocketMqMigrationTaskRequestType) UnmarshalJSON(b []byte) error {
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

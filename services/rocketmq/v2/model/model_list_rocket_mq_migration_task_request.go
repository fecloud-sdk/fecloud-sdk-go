package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListRocketMqMigrationTaskRequest struct {
	InstanceId string `json:"instance_id"`

	Id *string `json:"id,omitempty"`

	Type *ListRocketMqMigrationTaskRequestType `json:"type,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o ListRocketMqMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketMqMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ListRocketMqMigrationTaskRequest", string(data)}, " ")
}

type ListRocketMqMigrationTaskRequestType struct {
	value string
}

type ListRocketMqMigrationTaskRequestTypeEnum struct {
	VHOST    ListRocketMqMigrationTaskRequestType
	EXCHANGE ListRocketMqMigrationTaskRequestType
	QUEUE    ListRocketMqMigrationTaskRequestType
	ALL      ListRocketMqMigrationTaskRequestType
}

func GetListRocketMqMigrationTaskRequestTypeEnum() ListRocketMqMigrationTaskRequestTypeEnum {
	return ListRocketMqMigrationTaskRequestTypeEnum{
		VHOST: ListRocketMqMigrationTaskRequestType{
			value: "vhost",
		},
		EXCHANGE: ListRocketMqMigrationTaskRequestType{
			value: "exchange",
		},
		QUEUE: ListRocketMqMigrationTaskRequestType{
			value: "queue",
		},
		ALL: ListRocketMqMigrationTaskRequestType{
			value: "all",
		},
	}
}

func (c ListRocketMqMigrationTaskRequestType) Value() string {
	return c.value
}

func (c ListRocketMqMigrationTaskRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRocketMqMigrationTaskRequestType) UnmarshalJSON(b []byte) error {
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

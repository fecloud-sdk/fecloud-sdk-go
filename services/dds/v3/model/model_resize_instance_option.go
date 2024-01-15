package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ResizeInstanceOption struct {
	TargetType *ResizeInstanceOptionTargetType `json:"target_type,omitempty"`

	TargetId string `json:"target_id"`

	TargetSpecCode string `json:"target_spec_code"`
}

func (o ResizeInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceOption struct{}"
	}

	return strings.Join([]string{"ResizeInstanceOption", string(data)}, " ")
}

type ResizeInstanceOptionTargetType struct {
	value string
}

type ResizeInstanceOptionTargetTypeEnum struct {
	MONGOS ResizeInstanceOptionTargetType
	SHARD  ResizeInstanceOptionTargetType
}

func GetResizeInstanceOptionTargetTypeEnum() ResizeInstanceOptionTargetTypeEnum {
	return ResizeInstanceOptionTargetTypeEnum{
		MONGOS: ResizeInstanceOptionTargetType{
			value: "mongos",
		},
		SHARD: ResizeInstanceOptionTargetType{
			value: "shard",
		},
	}
}

func (c ResizeInstanceOptionTargetType) Value() string {
	return c.value
}

func (c ResizeInstanceOptionTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeInstanceOptionTargetType) UnmarshalJSON(b []byte) error {
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

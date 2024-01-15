package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateInstanceFlavorOption struct {
	Type CreateInstanceFlavorOptionType `json:"type"`

	Num string `json:"num"`

	Storage *string `json:"storage,omitempty"`

	Size *string `json:"size,omitempty"`

	SpecCode string `json:"spec_code"`
}

func (o CreateInstanceFlavorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceFlavorOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceFlavorOption", string(data)}, " ")
}

type CreateInstanceFlavorOptionType struct {
	value string
}

type CreateInstanceFlavorOptionTypeEnum struct {
	MONGOS  CreateInstanceFlavorOptionType
	SHARD   CreateInstanceFlavorOptionType
	CONFIG  CreateInstanceFlavorOptionType
	REPLICA CreateInstanceFlavorOptionType
	SINGLE  CreateInstanceFlavorOptionType
}

func GetCreateInstanceFlavorOptionTypeEnum() CreateInstanceFlavorOptionTypeEnum {
	return CreateInstanceFlavorOptionTypeEnum{
		MONGOS: CreateInstanceFlavorOptionType{
			value: "mongos",
		},
		SHARD: CreateInstanceFlavorOptionType{
			value: "shard",
		},
		CONFIG: CreateInstanceFlavorOptionType{
			value: "config",
		},
		REPLICA: CreateInstanceFlavorOptionType{
			value: "replica",
		},
		SINGLE: CreateInstanceFlavorOptionType{
			value: "single",
		},
	}
}

func (c CreateInstanceFlavorOptionType) Value() string {
	return c.value
}

func (c CreateInstanceFlavorOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceFlavorOptionType) UnmarshalJSON(b []byte) error {
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

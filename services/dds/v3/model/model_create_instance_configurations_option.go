package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateInstanceConfigurationsOption struct {
	Type CreateInstanceConfigurationsOptionType `json:"type"`

	ConfigurationId string `json:"configuration_id"`
}

func (o CreateInstanceConfigurationsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceConfigurationsOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceConfigurationsOption", string(data)}, " ")
}

type CreateInstanceConfigurationsOptionType struct {
	value string
}

type CreateInstanceConfigurationsOptionTypeEnum struct {
	MONGOS  CreateInstanceConfigurationsOptionType
	SHARD   CreateInstanceConfigurationsOptionType
	CONFIG  CreateInstanceConfigurationsOptionType
	REPLICA CreateInstanceConfigurationsOptionType
	SINGLE  CreateInstanceConfigurationsOptionType
}

func GetCreateInstanceConfigurationsOptionTypeEnum() CreateInstanceConfigurationsOptionTypeEnum {
	return CreateInstanceConfigurationsOptionTypeEnum{
		MONGOS: CreateInstanceConfigurationsOptionType{
			value: "mongos",
		},
		SHARD: CreateInstanceConfigurationsOptionType{
			value: "shard",
		},
		CONFIG: CreateInstanceConfigurationsOptionType{
			value: "config",
		},
		REPLICA: CreateInstanceConfigurationsOptionType{
			value: "replica",
		},
		SINGLE: CreateInstanceConfigurationsOptionType{
			value: "single",
		},
	}
}

func (c CreateInstanceConfigurationsOptionType) Value() string {
	return c.value
}

func (c CreateInstanceConfigurationsOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceConfigurationsOptionType) UnmarshalJSON(b []byte) error {
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

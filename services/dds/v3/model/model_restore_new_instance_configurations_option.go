package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type RestoreNewInstanceConfigurationsOption struct {
	Type RestoreNewInstanceConfigurationsOptionType `json:"type"`

	ConfigurationId string `json:"configuration_id"`
}

func (o RestoreNewInstanceConfigurationsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceConfigurationsOption struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceConfigurationsOption", string(data)}, " ")
}

type RestoreNewInstanceConfigurationsOptionType struct {
	value string
}

type RestoreNewInstanceConfigurationsOptionTypeEnum struct {
	MONGOS  RestoreNewInstanceConfigurationsOptionType
	SHARD   RestoreNewInstanceConfigurationsOptionType
	CONFIG  RestoreNewInstanceConfigurationsOptionType
	REPLICA RestoreNewInstanceConfigurationsOptionType
	SINGLE  RestoreNewInstanceConfigurationsOptionType
}

func GetRestoreNewInstanceConfigurationsOptionTypeEnum() RestoreNewInstanceConfigurationsOptionTypeEnum {
	return RestoreNewInstanceConfigurationsOptionTypeEnum{
		MONGOS: RestoreNewInstanceConfigurationsOptionType{
			value: "mongos",
		},
		SHARD: RestoreNewInstanceConfigurationsOptionType{
			value: "shard",
		},
		CONFIG: RestoreNewInstanceConfigurationsOptionType{
			value: "config",
		},
		REPLICA: RestoreNewInstanceConfigurationsOptionType{
			value: "replica",
		},
		SINGLE: RestoreNewInstanceConfigurationsOptionType{
			value: "single",
		},
	}
}

func (c RestoreNewInstanceConfigurationsOptionType) Value() string {
	return c.value
}

func (c RestoreNewInstanceConfigurationsOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreNewInstanceConfigurationsOptionType) UnmarshalJSON(b []byte) error {
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

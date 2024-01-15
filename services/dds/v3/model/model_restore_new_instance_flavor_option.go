package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type RestoreNewInstanceFlavorOption struct {
	Type RestoreNewInstanceFlavorOptionType `json:"type"`

	Num string `json:"num"`

	Size *string `json:"size,omitempty"`

	SpecCode string `json:"spec_code"`
}

func (o RestoreNewInstanceFlavorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceFlavorOption struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceFlavorOption", string(data)}, " ")
}

type RestoreNewInstanceFlavorOptionType struct {
	value string
}

type RestoreNewInstanceFlavorOptionTypeEnum struct {
	MONGOS  RestoreNewInstanceFlavorOptionType
	SHARD   RestoreNewInstanceFlavorOptionType
	CONFIG  RestoreNewInstanceFlavorOptionType
	REPLICA RestoreNewInstanceFlavorOptionType
	SINGLE  RestoreNewInstanceFlavorOptionType
}

func GetRestoreNewInstanceFlavorOptionTypeEnum() RestoreNewInstanceFlavorOptionTypeEnum {
	return RestoreNewInstanceFlavorOptionTypeEnum{
		MONGOS: RestoreNewInstanceFlavorOptionType{
			value: "mongos",
		},
		SHARD: RestoreNewInstanceFlavorOptionType{
			value: "shard",
		},
		CONFIG: RestoreNewInstanceFlavorOptionType{
			value: "config",
		},
		REPLICA: RestoreNewInstanceFlavorOptionType{
			value: "replica",
		},
		SINGLE: RestoreNewInstanceFlavorOptionType{
			value: "single",
		},
	}
}

func (c RestoreNewInstanceFlavorOptionType) Value() string {
	return c.value
}

func (c RestoreNewInstanceFlavorOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreNewInstanceFlavorOptionType) UnmarshalJSON(b []byte) error {
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

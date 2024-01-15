package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type RestartInstanceRequestBody struct {
	TargetType *RestartInstanceRequestBodyTargetType `json:"target_type,omitempty"`

	TargetId string `json:"target_id"`
}

func (o RestartInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestartInstanceRequestBody", string(data)}, " ")
}

type RestartInstanceRequestBodyTargetType struct {
	value string
}

type RestartInstanceRequestBodyTargetTypeEnum struct {
	MONGOS RestartInstanceRequestBodyTargetType
	SHARD  RestartInstanceRequestBodyTargetType
	CONFIG RestartInstanceRequestBodyTargetType
}

func GetRestartInstanceRequestBodyTargetTypeEnum() RestartInstanceRequestBodyTargetTypeEnum {
	return RestartInstanceRequestBodyTargetTypeEnum{
		MONGOS: RestartInstanceRequestBodyTargetType{
			value: "mongos",
		},
		SHARD: RestartInstanceRequestBodyTargetType{
			value: "shard",
		},
		CONFIG: RestartInstanceRequestBodyTargetType{
			value: "config",
		},
	}
}

func (c RestartInstanceRequestBodyTargetType) Value() string {
	return c.value
}

func (c RestartInstanceRequestBodyTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestartInstanceRequestBodyTargetType) UnmarshalJSON(b []byte) error {
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

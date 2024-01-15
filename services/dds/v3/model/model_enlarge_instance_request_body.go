package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type EnlargeInstanceRequestBody struct {
	Type EnlargeInstanceRequestBodyType `json:"type"`

	SpecCode string `json:"spec_code"`

	Num string `json:"num"`

	Volume *AddShardingNodeVolumeOption `json:"volume,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o EnlargeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"EnlargeInstanceRequestBody", string(data)}, " ")
}

type EnlargeInstanceRequestBodyType struct {
	value string
}

type EnlargeInstanceRequestBodyTypeEnum struct {
	MONGOS EnlargeInstanceRequestBodyType
	SHARD  EnlargeInstanceRequestBodyType
}

func GetEnlargeInstanceRequestBodyTypeEnum() EnlargeInstanceRequestBodyTypeEnum {
	return EnlargeInstanceRequestBodyTypeEnum{
		MONGOS: EnlargeInstanceRequestBodyType{
			value: "mongos",
		},
		SHARD: EnlargeInstanceRequestBodyType{
			value: "shard",
		},
	}
}

func (c EnlargeInstanceRequestBodyType) Value() string {
	return c.value
}

func (c EnlargeInstanceRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnlargeInstanceRequestBodyType) UnmarshalJSON(b []byte) error {
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

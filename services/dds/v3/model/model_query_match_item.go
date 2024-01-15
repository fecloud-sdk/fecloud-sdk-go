package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type QueryMatchItem struct {
	Key QueryMatchItemKey `json:"key"`

	Value string `json:"value"`
}

func (o QueryMatchItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryMatchItem struct{}"
	}

	return strings.Join([]string{"QueryMatchItem", string(data)}, " ")
}

type QueryMatchItemKey struct {
	value string
}

type QueryMatchItemKeyEnum struct {
	INSTANCE_NAME QueryMatchItemKey
	INSTANCE_ID   QueryMatchItemKey
}

func GetQueryMatchItemKeyEnum() QueryMatchItemKeyEnum {
	return QueryMatchItemKeyEnum{
		INSTANCE_NAME: QueryMatchItemKey{
			value: "instance_name",
		},
		INSTANCE_ID: QueryMatchItemKey{
			value: "instance_id",
		},
	}
}

func (c QueryMatchItemKey) Value() string {
	return c.value
}

func (c QueryMatchItemKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryMatchItemKey) UnmarshalJSON(b []byte) error {
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

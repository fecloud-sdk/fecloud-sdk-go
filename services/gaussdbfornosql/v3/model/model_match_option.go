package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type MatchOption struct {
	Key MatchOptionKey `json:"key"`

	Value string `json:"value"`
}

func (o MatchOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatchOption struct{}"
	}

	return strings.Join([]string{"MatchOption", string(data)}, " ")
}

type MatchOptionKey struct {
	value string
}

type MatchOptionKeyEnum struct {
	INSTANCE_NAME MatchOptionKey
	INSTANCE_ID   MatchOptionKey
}

func GetMatchOptionKeyEnum() MatchOptionKeyEnum {
	return MatchOptionKeyEnum{
		INSTANCE_NAME: MatchOptionKey{
			value: "instance_name",
		},
		INSTANCE_ID: MatchOptionKey{
			value: "instance_id",
		},
	}
}

func (c MatchOptionKey) Value() string {
	return c.value
}

func (c MatchOptionKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MatchOptionKey) UnmarshalJSON(b []byte) error {
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

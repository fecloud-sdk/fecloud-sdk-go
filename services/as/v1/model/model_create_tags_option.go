package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateTagsOption struct {
	Tags []TagsSingleValue `json:"tags"`

	Action CreateTagsOptionAction `json:"action"`
}

func (o CreateTagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsOption struct{}"
	}

	return strings.Join([]string{"CreateTagsOption", string(data)}, " ")
}

type CreateTagsOptionAction struct {
	value string
}

type CreateTagsOptionActionEnum struct {
	CREATE CreateTagsOptionAction
}

func GetCreateTagsOptionActionEnum() CreateTagsOptionActionEnum {
	return CreateTagsOptionActionEnum{
		CREATE: CreateTagsOptionAction{
			value: "create",
		},
	}
}

func (c CreateTagsOptionAction) Value() string {
	return c.value
}

func (c CreateTagsOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTagsOptionAction) UnmarshalJSON(b []byte) error {
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

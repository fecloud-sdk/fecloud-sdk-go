package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchAddSharedTagsRequestBody struct {
	Action BatchAddSharedTagsRequestBodyAction `json:"action"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`
}

func (o BatchAddSharedTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsRequestBody", string(data)}, " ")
}

type BatchAddSharedTagsRequestBodyAction struct {
	value string
}

type BatchAddSharedTagsRequestBodyActionEnum struct {
	CREATE BatchAddSharedTagsRequestBodyAction
}

func GetBatchAddSharedTagsRequestBodyActionEnum() BatchAddSharedTagsRequestBodyActionEnum {
	return BatchAddSharedTagsRequestBodyActionEnum{
		CREATE: BatchAddSharedTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchAddSharedTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchAddSharedTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddSharedTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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

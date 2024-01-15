package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchCreateListenerTagsRequestBody struct {
	Action BatchCreateListenerTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsRequestBody", string(data)}, " ")
}

type BatchCreateListenerTagsRequestBodyAction struct {
	value string
}

type BatchCreateListenerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateListenerTagsRequestBodyAction
}

func GetBatchCreateListenerTagsRequestBodyActionEnum() BatchCreateListenerTagsRequestBodyActionEnum {
	return BatchCreateListenerTagsRequestBodyActionEnum{
		CREATE: BatchCreateListenerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateListenerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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

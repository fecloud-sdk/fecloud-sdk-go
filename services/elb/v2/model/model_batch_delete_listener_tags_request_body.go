package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteListenerTagsRequestBody struct {
	Action BatchDeleteListenerTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteListenerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteListenerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteListenerTagsRequestBodyAction
}

func GetBatchDeleteListenerTagsRequestBodyActionEnum() BatchDeleteListenerTagsRequestBodyActionEnum {
	return BatchDeleteListenerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteListenerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteListenerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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

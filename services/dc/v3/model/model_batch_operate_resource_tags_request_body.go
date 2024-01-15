package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchOperateResourceTagsRequestBody struct {
	Action BatchOperateResourceTagsRequestBodyAction `json:"action"`

	Tags *[]Tag `json:"tags,omitempty"`

	SysTags *[]Tag `json:"sys_tags,omitempty"`
}

func (o BatchOperateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchOperateResourceTagsRequestBody", string(data)}, " ")
}

type BatchOperateResourceTagsRequestBodyAction struct {
	value string
}

type BatchOperateResourceTagsRequestBodyActionEnum struct {
	CREATE BatchOperateResourceTagsRequestBodyAction
	DELETE BatchOperateResourceTagsRequestBodyAction
}

func GetBatchOperateResourceTagsRequestBodyActionEnum() BatchOperateResourceTagsRequestBodyActionEnum {
	return BatchOperateResourceTagsRequestBodyActionEnum{
		CREATE: BatchOperateResourceTagsRequestBodyAction{
			value: "create",
		},
		DELETE: BatchOperateResourceTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchOperateResourceTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchOperateResourceTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchOperateResourceTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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

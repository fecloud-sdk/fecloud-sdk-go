package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchOperateInstanceTagRequestBody struct {
	Action BatchOperateInstanceTagRequestBodyAction `json:"action"`

	Tags []TagItem `json:"tags"`
}

func (o BatchOperateInstanceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateInstanceTagRequestBody struct{}"
	}

	return strings.Join([]string{"BatchOperateInstanceTagRequestBody", string(data)}, " ")
}

type BatchOperateInstanceTagRequestBodyAction struct {
	value string
}

type BatchOperateInstanceTagRequestBodyActionEnum struct {
	CREATE BatchOperateInstanceTagRequestBodyAction
	DELETE BatchOperateInstanceTagRequestBodyAction
}

func GetBatchOperateInstanceTagRequestBodyActionEnum() BatchOperateInstanceTagRequestBodyActionEnum {
	return BatchOperateInstanceTagRequestBodyActionEnum{
		CREATE: BatchOperateInstanceTagRequestBodyAction{
			value: "create",
		},
		DELETE: BatchOperateInstanceTagRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchOperateInstanceTagRequestBodyAction) Value() string {
	return c.value
}

func (c BatchOperateInstanceTagRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchOperateInstanceTagRequestBodyAction) UnmarshalJSON(b []byte) error {
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

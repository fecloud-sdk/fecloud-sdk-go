package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchTagActionRequestBody struct {
	Action BatchTagActionRequestBodyAction `json:"action"`

	Tags []BatchTagActionTagOption `json:"tags"`
}

func (o BatchTagActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequestBody", string(data)}, " ")
}

type BatchTagActionRequestBodyAction struct {
	value string
}

type BatchTagActionRequestBodyActionEnum struct {
	CREATE BatchTagActionRequestBodyAction
	DELETE BatchTagActionRequestBodyAction
}

func GetBatchTagActionRequestBodyActionEnum() BatchTagActionRequestBodyActionEnum {
	return BatchTagActionRequestBodyActionEnum{
		CREATE: BatchTagActionRequestBodyAction{
			value: "create",
		},
		DELETE: BatchTagActionRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchTagActionRequestBodyAction) Value() string {
	return c.value
}

func (c BatchTagActionRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchTagActionRequestBodyAction) UnmarshalJSON(b []byte) error {
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

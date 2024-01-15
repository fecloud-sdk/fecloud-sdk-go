package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchCreateOrDeleteTagReq struct {
	Action *BatchCreateOrDeleteTagReqAction `json:"action,omitempty"`

	Tags *[]TagEntity `json:"tags,omitempty"`
}

func (o BatchCreateOrDeleteTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagReq struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagReq", string(data)}, " ")
}

type BatchCreateOrDeleteTagReqAction struct {
	value string
}

type BatchCreateOrDeleteTagReqActionEnum struct {
	CREATE BatchCreateOrDeleteTagReqAction
	DELETE BatchCreateOrDeleteTagReqAction
}

func GetBatchCreateOrDeleteTagReqActionEnum() BatchCreateOrDeleteTagReqActionEnum {
	return BatchCreateOrDeleteTagReqActionEnum{
		CREATE: BatchCreateOrDeleteTagReqAction{
			value: "create",
		},
		DELETE: BatchCreateOrDeleteTagReqAction{
			value: "delete",
		},
	}
}

func (c BatchCreateOrDeleteTagReqAction) Value() string {
	return c.value
}

func (c BatchCreateOrDeleteTagReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateOrDeleteTagReqAction) UnmarshalJSON(b []byte) error {
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

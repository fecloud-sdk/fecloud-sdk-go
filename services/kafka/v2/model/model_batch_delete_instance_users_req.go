package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteInstanceUsersReq struct {
	Action *BatchDeleteInstanceUsersReqAction `json:"action,omitempty"`

	Users *[]string `json:"users,omitempty"`
}

func (o BatchDeleteInstanceUsersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceUsersReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceUsersReq", string(data)}, " ")
}

type BatchDeleteInstanceUsersReqAction struct {
	value string
}

type BatchDeleteInstanceUsersReqActionEnum struct {
	DELETE BatchDeleteInstanceUsersReqAction
}

func GetBatchDeleteInstanceUsersReqActionEnum() BatchDeleteInstanceUsersReqActionEnum {
	return BatchDeleteInstanceUsersReqActionEnum{
		DELETE: BatchDeleteInstanceUsersReqAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteInstanceUsersReqAction) Value() string {
	return c.value
}

func (c BatchDeleteInstanceUsersReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceUsersReqAction) UnmarshalJSON(b []byte) error {
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

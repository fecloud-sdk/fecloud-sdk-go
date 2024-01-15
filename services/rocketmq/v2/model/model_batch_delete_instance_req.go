package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteInstanceReq struct {
	Instances *[]string `json:"instances,omitempty"`

	Action BatchDeleteInstanceReqAction `json:"action"`

	AllFailure *BatchDeleteInstanceReqAllFailure `json:"allFailure,omitempty"`
}

func (o BatchDeleteInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceReq", string(data)}, " ")
}

type BatchDeleteInstanceReqAction struct {
	value string
}

type BatchDeleteInstanceReqActionEnum struct {
	DELETE BatchDeleteInstanceReqAction
}

func GetBatchDeleteInstanceReqActionEnum() BatchDeleteInstanceReqActionEnum {
	return BatchDeleteInstanceReqActionEnum{
		DELETE: BatchDeleteInstanceReqAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteInstanceReqAction) Value() string {
	return c.value
}

func (c BatchDeleteInstanceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceReqAction) UnmarshalJSON(b []byte) error {
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

type BatchDeleteInstanceReqAllFailure struct {
	value string
}

type BatchDeleteInstanceReqAllFailureEnum struct {
	RELIABILITY BatchDeleteInstanceReqAllFailure
}

func GetBatchDeleteInstanceReqAllFailureEnum() BatchDeleteInstanceReqAllFailureEnum {
	return BatchDeleteInstanceReqAllFailureEnum{
		RELIABILITY: BatchDeleteInstanceReqAllFailure{
			value: "reliability",
		},
	}
}

func (c BatchDeleteInstanceReqAllFailure) Value() string {
	return c.value
}

func (c BatchDeleteInstanceReqAllFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceReqAllFailure) UnmarshalJSON(b []byte) error {
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

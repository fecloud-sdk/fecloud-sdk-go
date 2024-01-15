package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchRestartOrDeleteInstanceReq struct {
	Instances *[]string `json:"instances,omitempty"`

	Action BatchRestartOrDeleteInstanceReqAction `json:"action"`

	AllFailure *BatchRestartOrDeleteInstanceReqAllFailure `json:"allFailure,omitempty"`
}

func (o BatchRestartOrDeleteInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOrDeleteInstanceReq struct{}"
	}

	return strings.Join([]string{"BatchRestartOrDeleteInstanceReq", string(data)}, " ")
}

type BatchRestartOrDeleteInstanceReqAction struct {
	value string
}

type BatchRestartOrDeleteInstanceReqActionEnum struct {
	RESTART BatchRestartOrDeleteInstanceReqAction
	DELETE  BatchRestartOrDeleteInstanceReqAction
}

func GetBatchRestartOrDeleteInstanceReqActionEnum() BatchRestartOrDeleteInstanceReqActionEnum {
	return BatchRestartOrDeleteInstanceReqActionEnum{
		RESTART: BatchRestartOrDeleteInstanceReqAction{
			value: "restart",
		},
		DELETE: BatchRestartOrDeleteInstanceReqAction{
			value: "delete",
		},
	}
}

func (c BatchRestartOrDeleteInstanceReqAction) Value() string {
	return c.value
}

func (c BatchRestartOrDeleteInstanceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestartOrDeleteInstanceReqAction) UnmarshalJSON(b []byte) error {
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

type BatchRestartOrDeleteInstanceReqAllFailure struct {
	value string
}

type BatchRestartOrDeleteInstanceReqAllFailureEnum struct {
	KAFKA BatchRestartOrDeleteInstanceReqAllFailure
}

func GetBatchRestartOrDeleteInstanceReqAllFailureEnum() BatchRestartOrDeleteInstanceReqAllFailureEnum {
	return BatchRestartOrDeleteInstanceReqAllFailureEnum{
		KAFKA: BatchRestartOrDeleteInstanceReqAllFailure{
			value: "kafka",
		},
	}
}

func (c BatchRestartOrDeleteInstanceReqAllFailure) Value() string {
	return c.value
}

func (c BatchRestartOrDeleteInstanceReqAllFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestartOrDeleteInstanceReqAllFailure) UnmarshalJSON(b []byte) error {
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

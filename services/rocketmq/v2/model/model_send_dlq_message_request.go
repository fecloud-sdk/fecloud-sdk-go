package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type SendDlqMessageRequest struct {
	Engine SendDlqMessageRequestEngine `json:"engine"`

	InstanceId string `json:"instance_id"`

	Body *DeadletterResendReq `json:"body,omitempty"`
}

func (o SendDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"SendDlqMessageRequest", string(data)}, " ")
}

type SendDlqMessageRequestEngine struct {
	value string
}

type SendDlqMessageRequestEngineEnum struct {
	RELIABILITY SendDlqMessageRequestEngine
}

func GetSendDlqMessageRequestEngineEnum() SendDlqMessageRequestEngineEnum {
	return SendDlqMessageRequestEngineEnum{
		RELIABILITY: SendDlqMessageRequestEngine{
			value: "reliability",
		},
	}
}

func (c SendDlqMessageRequestEngine) Value() string {
	return c.value
}

func (c SendDlqMessageRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SendDlqMessageRequestEngine) UnmarshalJSON(b []byte) error {
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

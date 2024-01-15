package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type AcceptOrRejectEndpointRequestBody struct {
	Action AcceptOrRejectEndpointRequestBodyAction `json:"action"`

	Endpoints []string `json:"endpoints"`
}

func (o AcceptOrRejectEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointRequestBody", string(data)}, " ")
}

type AcceptOrRejectEndpointRequestBodyAction struct {
	value string
}

type AcceptOrRejectEndpointRequestBodyActionEnum struct {
	RECEIVE AcceptOrRejectEndpointRequestBodyAction
	REJECT  AcceptOrRejectEndpointRequestBodyAction
}

func GetAcceptOrRejectEndpointRequestBodyActionEnum() AcceptOrRejectEndpointRequestBodyActionEnum {
	return AcceptOrRejectEndpointRequestBodyActionEnum{
		RECEIVE: AcceptOrRejectEndpointRequestBodyAction{
			value: "receive",
		},
		REJECT: AcceptOrRejectEndpointRequestBodyAction{
			value: "reject",
		},
	}
}

func (c AcceptOrRejectEndpointRequestBodyAction) Value() string {
	return c.value
}

func (c AcceptOrRejectEndpointRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptOrRejectEndpointRequestBodyAction) UnmarshalJSON(b []byte) error {
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

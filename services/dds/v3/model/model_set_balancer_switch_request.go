package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type SetBalancerSwitchRequest struct {
	InstanceId string `json:"instance_id"`

	Action SetBalancerSwitchRequestAction `json:"action"`
}

func (o SetBalancerSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBalancerSwitchRequest struct{}"
	}

	return strings.Join([]string{"SetBalancerSwitchRequest", string(data)}, " ")
}

type SetBalancerSwitchRequestAction struct {
	value string
}

type SetBalancerSwitchRequestActionEnum struct {
	START SetBalancerSwitchRequestAction
	STOP  SetBalancerSwitchRequestAction
}

func GetSetBalancerSwitchRequestActionEnum() SetBalancerSwitchRequestActionEnum {
	return SetBalancerSwitchRequestActionEnum{
		START: SetBalancerSwitchRequestAction{
			value: "start",
		},
		STOP: SetBalancerSwitchRequestAction{
			value: "stop",
		},
	}
}

func (c SetBalancerSwitchRequestAction) Value() string {
	return c.value
}

func (c SetBalancerSwitchRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetBalancerSwitchRequestAction) UnmarshalJSON(b []byte) error {
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

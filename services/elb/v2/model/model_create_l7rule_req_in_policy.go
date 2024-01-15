package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateL7ruleReqInPolicy struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Type CreateL7ruleReqInPolicyType `json:"type"`

	CompareType string `json:"compare_type"`

	Key *string `json:"key,omitempty"`

	Value string `json:"value"`

	Invert *bool `json:"invert,omitempty"`
}

func (o CreateL7ruleReqInPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleReqInPolicy struct{}"
	}

	return strings.Join([]string{"CreateL7ruleReqInPolicy", string(data)}, " ")
}

type CreateL7ruleReqInPolicyType struct {
	value string
}

type CreateL7ruleReqInPolicyTypeEnum struct {
	HOST_NAME CreateL7ruleReqInPolicyType
	PATH      CreateL7ruleReqInPolicyType
}

func GetCreateL7ruleReqInPolicyTypeEnum() CreateL7ruleReqInPolicyTypeEnum {
	return CreateL7ruleReqInPolicyTypeEnum{
		HOST_NAME: CreateL7ruleReqInPolicyType{
			value: "HOST_NAME",
		},
		PATH: CreateL7ruleReqInPolicyType{
			value: "PATH",
		},
	}
}

func (c CreateL7ruleReqInPolicyType) Value() string {
	return c.value
}

func (c CreateL7ruleReqInPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateL7ruleReqInPolicyType) UnmarshalJSON(b []byte) error {
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

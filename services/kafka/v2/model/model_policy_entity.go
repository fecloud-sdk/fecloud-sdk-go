package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type PolicyEntity struct {
	Owner *bool `json:"owner,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	AccessPolicy *PolicyEntityAccessPolicy `json:"access_policy,omitempty"`
}

func (o PolicyEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEntity struct{}"
	}

	return strings.Join([]string{"PolicyEntity", string(data)}, " ")
}

type PolicyEntityAccessPolicy struct {
	value string
}

type PolicyEntityAccessPolicyEnum struct {
	ALL PolicyEntityAccessPolicy
	PUB PolicyEntityAccessPolicy
	SUB PolicyEntityAccessPolicy
}

func GetPolicyEntityAccessPolicyEnum() PolicyEntityAccessPolicyEnum {
	return PolicyEntityAccessPolicyEnum{
		ALL: PolicyEntityAccessPolicy{
			value: "all",
		},
		PUB: PolicyEntityAccessPolicy{
			value: "pub",
		},
		SUB: PolicyEntityAccessPolicy{
			value: "sub",
		},
	}
}

func (c PolicyEntityAccessPolicy) Value() string {
	return c.value
}

func (c PolicyEntityAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyEntityAccessPolicy) UnmarshalJSON(b []byte) error {
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

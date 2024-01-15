package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type AccessPolicyEntity struct {
	UserName *string `json:"user_name,omitempty"`

	AccessPolicy *AccessPolicyEntityAccessPolicy `json:"access_policy,omitempty"`
}

func (o AccessPolicyEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyEntity struct{}"
	}

	return strings.Join([]string{"AccessPolicyEntity", string(data)}, " ")
}

type AccessPolicyEntityAccessPolicy struct {
	value string
}

type AccessPolicyEntityAccessPolicyEnum struct {
	ALL AccessPolicyEntityAccessPolicy
	PUB AccessPolicyEntityAccessPolicy
	SUB AccessPolicyEntityAccessPolicy
}

func GetAccessPolicyEntityAccessPolicyEnum() AccessPolicyEntityAccessPolicyEnum {
	return AccessPolicyEntityAccessPolicyEnum{
		ALL: AccessPolicyEntityAccessPolicy{
			value: "all",
		},
		PUB: AccessPolicyEntityAccessPolicy{
			value: "pub",
		},
		SUB: AccessPolicyEntityAccessPolicy{
			value: "sub",
		},
	}
}

func (c AccessPolicyEntityAccessPolicy) Value() string {
	return c.value
}

func (c AccessPolicyEntityAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyEntityAccessPolicy) UnmarshalJSON(b []byte) error {
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

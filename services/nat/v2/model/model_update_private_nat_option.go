package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdatePrivateNatOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Spec *UpdatePrivateNatOptionSpec `json:"spec,omitempty"`
}

func (o UpdatePrivateNatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateNatOption struct{}"
	}

	return strings.Join([]string{"UpdatePrivateNatOption", string(data)}, " ")
}

type UpdatePrivateNatOptionSpec struct {
	value string
}

type UpdatePrivateNatOptionSpecEnum struct {
	SMALL       UpdatePrivateNatOptionSpec
	MEDIUM      UpdatePrivateNatOptionSpec
	LARGE       UpdatePrivateNatOptionSpec
	EXTRA_LARGE UpdatePrivateNatOptionSpec
}

func GetUpdatePrivateNatOptionSpecEnum() UpdatePrivateNatOptionSpecEnum {
	return UpdatePrivateNatOptionSpecEnum{
		SMALL: UpdatePrivateNatOptionSpec{
			value: "Small",
		},
		MEDIUM: UpdatePrivateNatOptionSpec{
			value: "Medium",
		},
		LARGE: UpdatePrivateNatOptionSpec{
			value: "Large",
		},
		EXTRA_LARGE: UpdatePrivateNatOptionSpec{
			value: "Extra-large",
		},
	}
}

func (c UpdatePrivateNatOptionSpec) Value() string {
	return c.value
}

func (c UpdatePrivateNatOptionSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePrivateNatOptionSpec) UnmarshalJSON(b []byte) error {
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

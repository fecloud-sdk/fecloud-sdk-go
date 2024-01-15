package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreatePrivateNatOption struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Spec *CreatePrivateNatOptionSpec `json:"spec,omitempty"`

	DownlinkVpcs []DownlinkVpcOption `json:"downlink_vpcs"`

	Tags *[]PrivateTag `json:"tags,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePrivateNatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatOption", string(data)}, " ")
}

type CreatePrivateNatOptionSpec struct {
	value string
}

type CreatePrivateNatOptionSpecEnum struct {
	SMALL       CreatePrivateNatOptionSpec
	MEDIUM      CreatePrivateNatOptionSpec
	LARGE       CreatePrivateNatOptionSpec
	EXTRA_LARGE CreatePrivateNatOptionSpec
}

func GetCreatePrivateNatOptionSpecEnum() CreatePrivateNatOptionSpecEnum {
	return CreatePrivateNatOptionSpecEnum{
		SMALL: CreatePrivateNatOptionSpec{
			value: "Small",
		},
		MEDIUM: CreatePrivateNatOptionSpec{
			value: "Medium",
		},
		LARGE: CreatePrivateNatOptionSpec{
			value: "Large",
		},
		EXTRA_LARGE: CreatePrivateNatOptionSpec{
			value: "Extra-large",
		},
	}
}

func (c CreatePrivateNatOptionSpec) Value() string {
	return c.value
}

func (c CreatePrivateNatOptionSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivateNatOptionSpec) UnmarshalJSON(b []byte) error {
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

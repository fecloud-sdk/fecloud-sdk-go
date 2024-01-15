package model

import (
	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"
	"strings"
)

type PrivateNat struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Spec PrivateNatSpec `json:"spec"`

	Status PrivateNatStatus `json:"status"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	DownlinkVpcs []DownlinkVpc `json:"downlink_vpcs"`

	Tags *[]PrivateTag `json:"tags,omitempty"`

	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o PrivateNat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateNat struct{}"
	}

	return strings.Join([]string{"PrivateNat", string(data)}, " ")
}

type PrivateNatSpec struct {
	value string
}

type PrivateNatSpecEnum struct {
	SMALL       PrivateNatSpec
	MEDIUM      PrivateNatSpec
	LARGE       PrivateNatSpec
	EXTRA_LARGE PrivateNatSpec
}

func GetPrivateNatSpecEnum() PrivateNatSpecEnum {
	return PrivateNatSpecEnum{
		SMALL: PrivateNatSpec{
			value: "Small",
		},
		MEDIUM: PrivateNatSpec{
			value: "Medium",
		},
		LARGE: PrivateNatSpec{
			value: "Large",
		},
		EXTRA_LARGE: PrivateNatSpec{
			value: "Extra-large",
		},
	}
}

func (c PrivateNatSpec) Value() string {
	return c.value
}

func (c PrivateNatSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateNatSpec) UnmarshalJSON(b []byte) error {
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

type PrivateNatStatus struct {
	value string
}

type PrivateNatStatusEnum struct {
	ACTIVE PrivateNatStatus
	FROZEN PrivateNatStatus
}

func GetPrivateNatStatusEnum() PrivateNatStatusEnum {
	return PrivateNatStatusEnum{
		ACTIVE: PrivateNatStatus{
			value: "ACTIVE",
		},
		FROZEN: PrivateNatStatus{
			value: "FROZEN",
		},
	}
}

func (c PrivateNatStatus) Value() string {
	return c.value
}

func (c PrivateNatStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateNatStatus) UnmarshalJSON(b []byte) error {
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

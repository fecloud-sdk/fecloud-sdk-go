package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateL7ruleReq struct {
	TenantId *string `json:"tenant_id,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Type CreateL7ruleReqType `json:"type"`

	CompareType string `json:"compare_type"`

	Key *string `json:"key,omitempty"`

	Value string `json:"value"`

	Invert *bool `json:"invert,omitempty"`
}

func (o CreateL7ruleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleReq struct{}"
	}

	return strings.Join([]string{"CreateL7ruleReq", string(data)}, " ")
}

type CreateL7ruleReqType struct {
	value string
}

type CreateL7ruleReqTypeEnum struct {
	HOST_NAME CreateL7ruleReqType
	PATH      CreateL7ruleReqType
}

func GetCreateL7ruleReqTypeEnum() CreateL7ruleReqTypeEnum {
	return CreateL7ruleReqTypeEnum{
		HOST_NAME: CreateL7ruleReqType{
			value: "HOST_NAME",
		},
		PATH: CreateL7ruleReqType{
			value: "PATH",
		},
	}
}

func (c CreateL7ruleReqType) Value() string {
	return c.value
}

func (c CreateL7ruleReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateL7ruleReqType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type Quotas struct {
	Type *QuotasType `json:"type,omitempty"`

	Used *int32 `json:"used,omitempty"`

	Quota *int32 `json:"quota,omitempty"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}

type QuotasType struct {
	value string
}

type QuotasTypeEnum struct {
	ENDPOINT_SERVICE QuotasType
	ENDPOINT         QuotasType
}

func GetQuotasTypeEnum() QuotasTypeEnum {
	return QuotasTypeEnum{
		ENDPOINT_SERVICE: QuotasType{
			value: "endpoint_service",
		},
		ENDPOINT: QuotasType{
			value: "endpoint",
		},
	}
}

func (c QuotasType) Value() string {
	return c.value
}

func (c QuotasType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotasType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateCompareTaskRequest struct {
	XLanguage *CreateCompareTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateCompareTaskReq `json:"body,omitempty"`
}

func (o CreateCompareTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskRequest", string(data)}, " ")
}

type CreateCompareTaskRequestXLanguage struct {
	value string
}

type CreateCompareTaskRequestXLanguageEnum struct {
	EN_US CreateCompareTaskRequestXLanguage
	ZH_CN CreateCompareTaskRequestXLanguage
}

func GetCreateCompareTaskRequestXLanguageEnum() CreateCompareTaskRequestXLanguageEnum {
	return CreateCompareTaskRequestXLanguageEnum{
		EN_US: CreateCompareTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateCompareTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateCompareTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CreateCompareTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCompareTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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

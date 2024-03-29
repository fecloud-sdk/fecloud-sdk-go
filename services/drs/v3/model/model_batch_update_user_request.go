package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchUpdateUserRequest struct {
	XLanguage *BatchUpdateUserRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchUpdateSrcUserReq `json:"body,omitempty"`
}

func (o BatchUpdateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUserRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateUserRequest", string(data)}, " ")
}

type BatchUpdateUserRequestXLanguage struct {
	value string
}

type BatchUpdateUserRequestXLanguageEnum struct {
	EN_US BatchUpdateUserRequestXLanguage
	ZH_CN BatchUpdateUserRequestXLanguage
}

func GetBatchUpdateUserRequestXLanguageEnum() BatchUpdateUserRequestXLanguageEnum {
	return BatchUpdateUserRequestXLanguageEnum{
		EN_US: BatchUpdateUserRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchUpdateUserRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchUpdateUserRequestXLanguage) Value() string {
	return c.value
}

func (c BatchUpdateUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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

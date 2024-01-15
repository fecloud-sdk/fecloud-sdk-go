package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListErrorlogForLtsRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *ListErrorlogForLtsRequestXLanguage `json:"X-Language,omitempty"`

	Body *ErrorlogForLtsRequest `json:"body,omitempty"`
}

func (o ListErrorlogForLtsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorlogForLtsRequest struct{}"
	}

	return strings.Join([]string{"ListErrorlogForLtsRequest", string(data)}, " ")
}

type ListErrorlogForLtsRequestXLanguage struct {
	value string
}

type ListErrorlogForLtsRequestXLanguageEnum struct {
	ZH_CN ListErrorlogForLtsRequestXLanguage
	EN_US ListErrorlogForLtsRequestXLanguage
}

func GetListErrorlogForLtsRequestXLanguageEnum() ListErrorlogForLtsRequestXLanguageEnum {
	return ListErrorlogForLtsRequestXLanguageEnum{
		ZH_CN: ListErrorlogForLtsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListErrorlogForLtsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListErrorlogForLtsRequestXLanguage) Value() string {
	return c.value
}

func (c ListErrorlogForLtsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListErrorlogForLtsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

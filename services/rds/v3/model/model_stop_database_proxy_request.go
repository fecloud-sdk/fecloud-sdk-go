package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type StopDatabaseProxyRequest struct {
	XLanguage *StopDatabaseProxyRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o StopDatabaseProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDatabaseProxyRequest struct{}"
	}

	return strings.Join([]string{"StopDatabaseProxyRequest", string(data)}, " ")
}

type StopDatabaseProxyRequestXLanguage struct {
	value string
}

type StopDatabaseProxyRequestXLanguageEnum struct {
	ZH_CN StopDatabaseProxyRequestXLanguage
	EN_US StopDatabaseProxyRequestXLanguage
}

func GetStopDatabaseProxyRequestXLanguageEnum() StopDatabaseProxyRequestXLanguageEnum {
	return StopDatabaseProxyRequestXLanguageEnum{
		ZH_CN: StopDatabaseProxyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopDatabaseProxyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopDatabaseProxyRequestXLanguage) Value() string {
	return c.value
}

func (c StopDatabaseProxyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopDatabaseProxyRequestXLanguage) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type StartInstanceRestartActionRequest struct {
	XLanguage *StartInstanceRestartActionRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *InstanceRestartRequsetBody `json:"body,omitempty"`
}

func (o StartInstanceRestartActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceRestartActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceRestartActionRequest", string(data)}, " ")
}

type StartInstanceRestartActionRequestXLanguage struct {
	value string
}

type StartInstanceRestartActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceRestartActionRequestXLanguage
	EN_US StartInstanceRestartActionRequestXLanguage
}

func GetStartInstanceRestartActionRequestXLanguageEnum() StartInstanceRestartActionRequestXLanguageEnum {
	return StartInstanceRestartActionRequestXLanguageEnum{
		ZH_CN: StartInstanceRestartActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceRestartActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceRestartActionRequestXLanguage) Value() string {
	return c.value
}

func (c StartInstanceRestartActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartInstanceRestartActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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

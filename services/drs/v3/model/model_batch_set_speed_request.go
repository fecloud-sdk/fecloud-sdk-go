package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchSetSpeedRequest struct {
	XLanguage *BatchSetSpeedRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchLimitSpeedReq `json:"body,omitempty"`
}

func (o BatchSetSpeedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSpeedRequest struct{}"
	}

	return strings.Join([]string{"BatchSetSpeedRequest", string(data)}, " ")
}

type BatchSetSpeedRequestXLanguage struct {
	value string
}

type BatchSetSpeedRequestXLanguageEnum struct {
	EN_US BatchSetSpeedRequestXLanguage
	ZH_CN BatchSetSpeedRequestXLanguage
}

func GetBatchSetSpeedRequestXLanguageEnum() BatchSetSpeedRequestXLanguageEnum {
	return BatchSetSpeedRequestXLanguageEnum{
		EN_US: BatchSetSpeedRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetSpeedRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetSpeedRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetSpeedRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetSpeedRequestXLanguage) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListUsersRequest struct {
	XLanguage *ListUsersRequestXLanguage `json:"X-Language,omitempty"`

	JobId string `json:"job_id"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}

type ListUsersRequestXLanguage struct {
	value string
}

type ListUsersRequestXLanguageEnum struct {
	EN_US ListUsersRequestXLanguage
	ZH_CN ListUsersRequestXLanguage
}

func GetListUsersRequestXLanguageEnum() ListUsersRequestXLanguageEnum {
	return ListUsersRequestXLanguageEnum{
		EN_US: ListUsersRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListUsersRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListUsersRequestXLanguage) Value() string {
	return c.value
}

func (c ListUsersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsersRequestXLanguage) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListQuotaDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	Type *ListQuotaDetailsRequestType `json:"type,omitempty"`
}

func (o ListQuotaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsRequest", string(data)}, " ")
}

type ListQuotaDetailsRequestType struct {
	value string
}

type ListQuotaDetailsRequestTypeEnum struct {
	ENDPOINT_SERVICE ListQuotaDetailsRequestType
	ENDPOINT         ListQuotaDetailsRequestType
}

func GetListQuotaDetailsRequestTypeEnum() ListQuotaDetailsRequestTypeEnum {
	return ListQuotaDetailsRequestTypeEnum{
		ENDPOINT_SERVICE: ListQuotaDetailsRequestType{
			value: "endpoint_service",
		},
		ENDPOINT: ListQuotaDetailsRequestType{
			value: "endpoint",
		},
	}
}

func (c ListQuotaDetailsRequestType) Value() string {
	return c.value
}

func (c ListQuotaDetailsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQuotaDetailsRequestType) UnmarshalJSON(b []byte) error {
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

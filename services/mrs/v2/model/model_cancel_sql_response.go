package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CancelSqlResponse struct {
	Message *string `json:"message,omitempty"`

	Status         *CancelSqlResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CancelSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSqlResponse struct{}"
	}

	return strings.Join([]string{"CancelSqlResponse", string(data)}, " ")
}

type CancelSqlResponseStatus struct {
	value string
}

type CancelSqlResponseStatusEnum struct {
	SUCCEED CancelSqlResponseStatus
	FAILED  CancelSqlResponseStatus
}

func GetCancelSqlResponseStatusEnum() CancelSqlResponseStatusEnum {
	return CancelSqlResponseStatusEnum{
		SUCCEED: CancelSqlResponseStatus{
			value: "SUCCEED",
		},
		FAILED: CancelSqlResponseStatus{
			value: "FAILED",
		},
	}
}

func (c CancelSqlResponseStatus) Value() string {
	return c.value
}

func (c CancelSqlResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelSqlResponseStatus) UnmarshalJSON(b []byte) error {
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

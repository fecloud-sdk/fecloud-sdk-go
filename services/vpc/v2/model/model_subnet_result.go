package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type SubnetResult struct {
	Id string `json:"id"`

	Status SubnetResultStatus `json:"status"`
}

func (o SubnetResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetResult struct{}"
	}

	return strings.Join([]string{"SubnetResult", string(data)}, " ")
}

type SubnetResultStatus struct {
	value string
}

type SubnetResultStatusEnum struct {
	ACTIVE  SubnetResultStatus
	UNKNOWN SubnetResultStatus
	ERROR   SubnetResultStatus
}

func GetSubnetResultStatusEnum() SubnetResultStatusEnum {
	return SubnetResultStatusEnum{
		ACTIVE: SubnetResultStatus{
			value: "ACTIVE",
		},
		UNKNOWN: SubnetResultStatus{
			value: "UNKNOWN",
		},
		ERROR: SubnetResultStatus{
			value: "ERROR",
		},
	}
}

func (c SubnetResultStatus) Value() string {
	return c.value
}

func (c SubnetResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubnetResultStatus) UnmarshalJSON(b []byte) error {
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

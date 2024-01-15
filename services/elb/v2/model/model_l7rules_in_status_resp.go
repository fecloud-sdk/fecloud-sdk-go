package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type L7rulesInStatusResp struct {
	Type L7rulesInStatusRespType `json:"type"`

	Id string `json:"id"`

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o L7rulesInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7rulesInStatusResp struct{}"
	}

	return strings.Join([]string{"L7rulesInStatusResp", string(data)}, " ")
}

type L7rulesInStatusRespType struct {
	value string
}

type L7rulesInStatusRespTypeEnum struct {
	PATH      L7rulesInStatusRespType
	HOST_NAME L7rulesInStatusRespType
}

func GetL7rulesInStatusRespTypeEnum() L7rulesInStatusRespTypeEnum {
	return L7rulesInStatusRespTypeEnum{
		PATH: L7rulesInStatusRespType{
			value: "PATH",
		},
		HOST_NAME: L7rulesInStatusRespType{
			value: "HOST_NAME",
		},
	}
}

func (c L7rulesInStatusRespType) Value() string {
	return c.value
}

func (c L7rulesInStatusRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7rulesInStatusRespType) UnmarshalJSON(b []byte) error {
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

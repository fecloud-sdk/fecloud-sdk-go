package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type HealthmonitorsInStatusResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Type HealthmonitorsInStatusRespType `json:"type"`

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o HealthmonitorsInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthmonitorsInStatusResp struct{}"
	}

	return strings.Join([]string{"HealthmonitorsInStatusResp", string(data)}, " ")
}

type HealthmonitorsInStatusRespType struct {
	value string
}

type HealthmonitorsInStatusRespTypeEnum struct {
	UDP_CONNECT HealthmonitorsInStatusRespType
	TCP         HealthmonitorsInStatusRespType
	HTTP        HealthmonitorsInStatusRespType
}

func GetHealthmonitorsInStatusRespTypeEnum() HealthmonitorsInStatusRespTypeEnum {
	return HealthmonitorsInStatusRespTypeEnum{
		UDP_CONNECT: HealthmonitorsInStatusRespType{
			value: "UDP_CONNECT",
		},
		TCP: HealthmonitorsInStatusRespType{
			value: "TCP",
		},
		HTTP: HealthmonitorsInStatusRespType{
			value: "HTTP",
		},
	}
}

func (c HealthmonitorsInStatusRespType) Value() string {
	return c.value
}

func (c HealthmonitorsInStatusRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthmonitorsInStatusRespType) UnmarshalJSON(b []byte) error {
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

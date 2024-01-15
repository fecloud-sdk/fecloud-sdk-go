package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type L7policiesInStatusResp struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Rules []L7rulesInStatusResp `json:"rules"`

	Action L7policiesInStatusRespAction `json:"action"`

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o L7policiesInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7policiesInStatusResp struct{}"
	}

	return strings.Join([]string{"L7policiesInStatusResp", string(data)}, " ")
}

type L7policiesInStatusRespAction struct {
	value string
}

type L7policiesInStatusRespActionEnum struct {
	REDIRECT_TO_POOL     L7policiesInStatusRespAction
	REDIRECT_TO_LISTENER L7policiesInStatusRespAction
}

func GetL7policiesInStatusRespActionEnum() L7policiesInStatusRespActionEnum {
	return L7policiesInStatusRespActionEnum{
		REDIRECT_TO_POOL: L7policiesInStatusRespAction{
			value: "REDIRECT_TO_POOL",
		},
		REDIRECT_TO_LISTENER: L7policiesInStatusRespAction{
			value: "REDIRECT_TO_LISTENER",
		},
	}
}

func (c L7policiesInStatusRespAction) Value() string {
	return c.value
}

func (c L7policiesInStatusRespAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7policiesInStatusRespAction) UnmarshalJSON(b []byte) error {
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

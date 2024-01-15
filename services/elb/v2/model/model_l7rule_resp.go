package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type L7ruleResp struct {
	Id string `json:"id"`

	ProvisioningStatus string `json:"provisioning_status"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`

	AdminStateUp bool `json:"admin_state_up"`

	Type L7ruleRespType `json:"type"`

	CompareType string `json:"compare_type"`

	Invert bool `json:"invert"`

	Key string `json:"key"`

	Value string `json:"value"`
}

func (o L7ruleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7ruleResp struct{}"
	}

	return strings.Join([]string{"L7ruleResp", string(data)}, " ")
}

type L7ruleRespType struct {
	value string
}

type L7ruleRespTypeEnum struct {
	HOST_NAME L7ruleRespType
	PATH      L7ruleRespType
}

func GetL7ruleRespTypeEnum() L7ruleRespTypeEnum {
	return L7ruleRespTypeEnum{
		HOST_NAME: L7ruleRespType{
			value: "HOST_NAME",
		},
		PATH: L7ruleRespType{
			value: "PATH",
		},
	}
}

func (c L7ruleRespType) Value() string {
	return c.value
}

func (c L7ruleRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7ruleRespType) UnmarshalJSON(b []byte) error {
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

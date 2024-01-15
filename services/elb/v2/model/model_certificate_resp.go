package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CertificateResp struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	AdminStateUp bool `json:"admin_state_up"`

	Name string `json:"name"`

	Description string `json:"description"`

	Type CertificateRespType `json:"type"`

	Domain string `json:"domain"`

	PrivateKey string `json:"private_key"`

	Certificate string `json:"certificate"`

	ExpireTime string `json:"expire_time"`

	CreateTime string `json:"create_time"`

	UpdateTime string `json:"update_time"`
}

func (o CertificateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateResp struct{}"
	}

	return strings.Join([]string{"CertificateResp", string(data)}, " ")
}

type CertificateRespType struct {
	value string
}

type CertificateRespTypeEnum struct {
	SERVER CertificateRespType
	CLIENT CertificateRespType
}

func GetCertificateRespTypeEnum() CertificateRespTypeEnum {
	return CertificateRespTypeEnum{
		SERVER: CertificateRespType{
			value: "server",
		},
		CLIENT: CertificateRespType{
			value: "client",
		},
	}
}

func (c CertificateRespType) Value() string {
	return c.value
}

func (c CertificateRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateRespType) UnmarshalJSON(b []byte) error {
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

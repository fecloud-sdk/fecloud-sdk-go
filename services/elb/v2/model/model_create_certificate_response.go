package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateCertificateResponse struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *CreateCertificateResponseType `json:"type,omitempty"`

	Domain *string `json:"domain,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	ExpireTime *string `json:"expire_time,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateResponse", string(data)}, " ")
}

type CreateCertificateResponseType struct {
	value string
}

type CreateCertificateResponseTypeEnum struct {
	SERVER CreateCertificateResponseType
	CLIENT CreateCertificateResponseType
}

func GetCreateCertificateResponseTypeEnum() CreateCertificateResponseTypeEnum {
	return CreateCertificateResponseTypeEnum{
		SERVER: CreateCertificateResponseType{
			value: "server",
		},
		CLIENT: CreateCertificateResponseType{
			value: "client",
		},
	}
}

func (c CreateCertificateResponseType) Value() string {
	return c.value
}

func (c CreateCertificateResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCertificateResponseType) UnmarshalJSON(b []byte) error {
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

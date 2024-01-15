package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateCertificateResponse struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *UpdateCertificateResponseType `json:"type,omitempty"`

	Domain *string `json:"domain,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	ExpireTime *string `json:"expire_time,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCertificateResponse", string(data)}, " ")
}

type UpdateCertificateResponseType struct {
	value string
}

type UpdateCertificateResponseTypeEnum struct {
	SERVER UpdateCertificateResponseType
	CLIENT UpdateCertificateResponseType
}

func GetUpdateCertificateResponseTypeEnum() UpdateCertificateResponseTypeEnum {
	return UpdateCertificateResponseTypeEnum{
		SERVER: UpdateCertificateResponseType{
			value: "server",
		},
		CLIENT: UpdateCertificateResponseType{
			value: "client",
		},
	}
}

func (c UpdateCertificateResponseType) Value() string {
	return c.value
}

func (c UpdateCertificateResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateResponseType) UnmarshalJSON(b []byte) error {
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

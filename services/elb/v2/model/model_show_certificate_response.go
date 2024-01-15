package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowCertificateResponse struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *ShowCertificateResponseType `json:"type,omitempty"`

	Domain *string `json:"domain,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	ExpireTime *string `json:"expire_time,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateResponse", string(data)}, " ")
}

type ShowCertificateResponseType struct {
	value string
}

type ShowCertificateResponseTypeEnum struct {
	SERVER ShowCertificateResponseType
	CLIENT ShowCertificateResponseType
}

func GetShowCertificateResponseTypeEnum() ShowCertificateResponseTypeEnum {
	return ShowCertificateResponseTypeEnum{
		SERVER: ShowCertificateResponseType{
			value: "server",
		},
		CLIENT: ShowCertificateResponseType{
			value: "client",
		},
	}
}

func (c ShowCertificateResponseType) Value() string {
	return c.value
}

func (c ShowCertificateResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCertificateResponseType) UnmarshalJSON(b []byte) error {
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

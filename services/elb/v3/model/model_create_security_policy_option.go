package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateSecurityPolicyOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Protocols []string `json:"protocols"`

	Ciphers []CreateSecurityPolicyOptionCiphers `json:"ciphers"`
}

func (o CreateSecurityPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyOption", string(data)}, " ")
}

type CreateSecurityPolicyOptionCiphers struct {
	value string
}

type CreateSecurityPolicyOptionCiphersEnum struct {
	ECDHE_RSA_AES256_GCM_SHA384   CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_GCM_SHA256   CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_GCM_SHA384 CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_GCM_SHA256 CreateSecurityPolicyOptionCiphers
	AES128_GCM_SHA256             CreateSecurityPolicyOptionCiphers
	AES256_GCM_SHA384             CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA256     CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA256       CreateSecurityPolicyOptionCiphers
	AES128_SHA256                 CreateSecurityPolicyOptionCiphers
	AES256_SHA256                 CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA384     CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA384       CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA        CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA          CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA          CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA        CreateSecurityPolicyOptionCiphers
	AES128_SHA                    CreateSecurityPolicyOptionCiphers
	AES256_SHA                    CreateSecurityPolicyOptionCiphers
	CAMELLIA128_SHA               CreateSecurityPolicyOptionCiphers
	DES_CBC3_SHA                  CreateSecurityPolicyOptionCiphers
	CAMELLIA256_SHA               CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_CHACHA20_POLY1305   CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_CHACHA20_POLY1305 CreateSecurityPolicyOptionCiphers
	TLS_AES_128_GCM_SHA256        CreateSecurityPolicyOptionCiphers
	TLS_AES_256_GCM_SHA384        CreateSecurityPolicyOptionCiphers
	TLS_CHACHA20_POLY1305_SHA256  CreateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_SHA256        CreateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_8_SHA256      CreateSecurityPolicyOptionCiphers
}

func GetCreateSecurityPolicyOptionCiphersEnum() CreateSecurityPolicyOptionCiphersEnum {
	return CreateSecurityPolicyOptionCiphersEnum{
		ECDHE_RSA_AES256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-GCM-SHA384",
		},
		ECDHE_RSA_AES128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-GCM-SHA256",
		},
		ECDHE_ECDSA_AES256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-GCM-SHA256",
		},
		AES128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "AES128-GCM-SHA256",
		},
		AES256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA256",
		},
		ECDHE_RSA_AES128_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA256",
		},
		AES128_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "AES128-SHA256",
		},
		AES256_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "AES256-SHA256",
		},
		ECDHE_ECDSA_AES256_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA384",
		},
		ECDHE_RSA_AES256_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA",
		},
		ECDHE_RSA_AES128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA",
		},
		ECDHE_RSA_AES256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA",
		},
		ECDHE_ECDSA_AES256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA",
		},
		AES128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "AES128-SHA",
		},
		AES256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "AES256-SHA",
		},
		CAMELLIA128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "CAMELLIA128-SHA",
		},
		DES_CBC3_SHA: CreateSecurityPolicyOptionCiphers{
			value: "DES-CBC3-SHA",
		},
		CAMELLIA256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "CAMELLIA256-SHA",
		},
		ECDHE_RSA_CHACHA20_POLY1305: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-CHACHA20-POLY1305",
		},
		ECDHE_ECDSA_CHACHA20_POLY1305: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-CHACHA20-POLY1305",
		},
		TLS_AES_128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_GCM_SHA256",
		},
		TLS_AES_256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_256_GCM_SHA384",
		},
		TLS_CHACHA20_POLY1305_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_CHACHA20_POLY1305_SHA256",
		},
		TLS_AES_128_CCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_SHA256",
		},
		TLS_AES_128_CCM_8_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_8_SHA256",
		},
	}
}

func (c CreateSecurityPolicyOptionCiphers) Value() string {
	return c.value
}

func (c CreateSecurityPolicyOptionCiphers) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPolicyOptionCiphers) UnmarshalJSON(b []byte) error {
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

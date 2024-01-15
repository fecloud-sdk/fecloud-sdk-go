package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateListenerReq struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Protocol CreateListenerReqProtocol `json:"protocol"`

	ProtocolPort int32 `json:"protocol_port"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ConnectionLimit *int32 `json:"connection_limit,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	InsertHeaders *InsertHeader `json:"insert_headers,omitempty"`

	TlsCiphersPolicy *CreateListenerReqTlsCiphersPolicy `json:"tls_ciphers_policy,omitempty"`
}

func (o CreateListenerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerReq struct{}"
	}

	return strings.Join([]string{"CreateListenerReq", string(data)}, " ")
}

type CreateListenerReqProtocol struct {
	value string
}

type CreateListenerReqProtocolEnum struct {
	UDP              CreateListenerReqProtocol
	TCP              CreateListenerReqProtocol
	HTTP             CreateListenerReqProtocol
	TERMINATED_HTTPS CreateListenerReqProtocol
}

func GetCreateListenerReqProtocolEnum() CreateListenerReqProtocolEnum {
	return CreateListenerReqProtocolEnum{
		UDP: CreateListenerReqProtocol{
			value: "UDP",
		},
		TCP: CreateListenerReqProtocol{
			value: "TCP",
		},
		HTTP: CreateListenerReqProtocol{
			value: "HTTP",
		},
		TERMINATED_HTTPS: CreateListenerReqProtocol{
			value: "TERMINATED_HTTPS",
		},
	}
}

func (c CreateListenerReqProtocol) Value() string {
	return c.value
}

func (c CreateListenerReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerReqProtocol) UnmarshalJSON(b []byte) error {
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

type CreateListenerReqTlsCiphersPolicy struct {
	value string
}

type CreateListenerReqTlsCiphersPolicyEnum struct {
	TLS_1_0        CreateListenerReqTlsCiphersPolicy
	TLS_1_1        CreateListenerReqTlsCiphersPolicy
	TLS_1_2        CreateListenerReqTlsCiphersPolicy
	TLS_1_2_STRICT CreateListenerReqTlsCiphersPolicy
}

func GetCreateListenerReqTlsCiphersPolicyEnum() CreateListenerReqTlsCiphersPolicyEnum {
	return CreateListenerReqTlsCiphersPolicyEnum{
		TLS_1_0: CreateListenerReqTlsCiphersPolicy{
			value: "tls-1-0",
		},
		TLS_1_1: CreateListenerReqTlsCiphersPolicy{
			value: "tls-1-1",
		},
		TLS_1_2: CreateListenerReqTlsCiphersPolicy{
			value: "tls-1-2",
		},
		TLS_1_2_STRICT: CreateListenerReqTlsCiphersPolicy{
			value: "tls-1-2-strict",
		},
	}
}

func (c CreateListenerReqTlsCiphersPolicy) Value() string {
	return c.value
}

func (c CreateListenerReqTlsCiphersPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerReqTlsCiphersPolicy) UnmarshalJSON(b []byte) error {
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

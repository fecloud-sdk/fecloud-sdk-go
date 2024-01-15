package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListenerResp struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	AdminStateUp bool `json:"admin_state_up"`

	Loadbalancers []ResourceList `json:"loadbalancers"`

	ConnectionLimit int32 `json:"connection_limit"`

	Http2Enable bool `json:"http2_enable"`

	Protocol ListenerRespProtocol `json:"protocol"`

	ProtocolPort int32 `json:"protocol_port"`

	DefaultPoolId string `json:"default_pool_id"`

	DefaultTlsContainerRef string `json:"default_tls_container_ref"`

	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`

	SniContainerRefs []string `json:"sni_container_refs"`

	Tags []string `json:"tags"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`

	InsertHeaders *InsertHeader `json:"insert_headers"`

	ProjectId string `json:"project_id"`

	TlsCiphersPolicy string `json:"tls_ciphers_policy"`

	ProtectionStatus ListenerRespProtectionStatus `json:"protection_status"`

	ProtectionReason string `json:"protection_reason"`
}

func (o ListenerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerResp struct{}"
	}

	return strings.Join([]string{"ListenerResp", string(data)}, " ")
}

type ListenerRespProtocol struct {
	value string
}

type ListenerRespProtocolEnum struct {
	UDP              ListenerRespProtocol
	TCP              ListenerRespProtocol
	HTTP             ListenerRespProtocol
	TERMINATED_HTTPS ListenerRespProtocol
}

func GetListenerRespProtocolEnum() ListenerRespProtocolEnum {
	return ListenerRespProtocolEnum{
		UDP: ListenerRespProtocol{
			value: "UDP",
		},
		TCP: ListenerRespProtocol{
			value: "TCP",
		},
		HTTP: ListenerRespProtocol{
			value: "HTTP",
		},
		TERMINATED_HTTPS: ListenerRespProtocol{
			value: "TERMINATED_HTTPS",
		},
	}
}

func (c ListenerRespProtocol) Value() string {
	return c.value
}

func (c ListenerRespProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerRespProtocol) UnmarshalJSON(b []byte) error {
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

type ListenerRespProtectionStatus struct {
	value string
}

type ListenerRespProtectionStatusEnum struct {
	NON_PROTECTION     ListenerRespProtectionStatus
	CONSOLE_PROTECTION ListenerRespProtectionStatus
}

func GetListenerRespProtectionStatusEnum() ListenerRespProtectionStatusEnum {
	return ListenerRespProtectionStatusEnum{
		NON_PROTECTION: ListenerRespProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: ListenerRespProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c ListenerRespProtectionStatus) Value() string {
	return c.value
}

func (c ListenerRespProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerRespProtectionStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateListenerOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	Description *string `json:"description,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`

	LoadbalancerId string `json:"loadbalancer_id"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Protocol string `json:"protocol"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	SniMatchAlgo *string `json:"sni_match_algo,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	SecurityPolicyId *string `json:"security_policy_id,omitempty"`

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	ClientTimeout *int32 `json:"client_timeout,omitempty"`

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	Ipgroup *CreateListenerIpGroupOption `json:"ipgroup,omitempty"`

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	QuicConfig *CreateListenerQuicConfigOption `json:"quic_config,omitempty"`

	ProtectionStatus *CreateListenerOptionProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`

	GzipEnable *bool `json:"gzip_enable,omitempty"`

	PortRanges *[]PortRange `json:"port_ranges,omitempty"`
}

func (o CreateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerOption struct{}"
	}

	return strings.Join([]string{"CreateListenerOption", string(data)}, " ")
}

type CreateListenerOptionProtectionStatus struct {
	value string
}

type CreateListenerOptionProtectionStatusEnum struct {
	NON_PROTECTION     CreateListenerOptionProtectionStatus
	CONSOLE_PROTECTION CreateListenerOptionProtectionStatus
}

func GetCreateListenerOptionProtectionStatusEnum() CreateListenerOptionProtectionStatusEnum {
	return CreateListenerOptionProtectionStatusEnum{
		NON_PROTECTION: CreateListenerOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateListenerOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateListenerOptionProtectionStatus) Value() string {
	return c.value
}

func (c CreateListenerOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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

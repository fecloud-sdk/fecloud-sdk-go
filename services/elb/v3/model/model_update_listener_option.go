package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateListenerOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	Description *string `json:"description,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`

	Name *string `json:"name,omitempty"`

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	SniMatchAlgo *string `json:"sni_match_algo,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	SecurityPolicyId *string `json:"security_policy_id,omitempty"`

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	ClientTimeout *int32 `json:"client_timeout,omitempty"`

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	Ipgroup *UpdateListenerIpGroupOption `json:"ipgroup,omitempty"`

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	QuicConfig *UpdateListenerQuicConfigOption `json:"quic_config,omitempty"`

	ProtectionStatus *UpdateListenerOptionProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`

	GzipEnable *bool `json:"gzip_enable,omitempty"`
}

func (o UpdateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerOption", string(data)}, " ")
}

type UpdateListenerOptionProtectionStatus struct {
	value string
}

type UpdateListenerOptionProtectionStatusEnum struct {
	NON_PROTECTION     UpdateListenerOptionProtectionStatus
	CONSOLE_PROTECTION UpdateListenerOptionProtectionStatus
}

func GetUpdateListenerOptionProtectionStatusEnum() UpdateListenerOptionProtectionStatusEnum {
	return UpdateListenerOptionProtectionStatusEnum{
		NON_PROTECTION: UpdateListenerOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateListenerOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateListenerOptionProtectionStatus) Value() string {
	return c.value
}

func (c UpdateListenerOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type Listener struct {
	AdminStateUp bool `json:"admin_state_up"`

	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`

	ConnectionLimit int32 `json:"connection_limit"`

	CreatedAt string `json:"created_at"`

	DefaultPoolId string `json:"default_pool_id"`

	DefaultTlsContainerRef string `json:"default_tls_container_ref"`

	Description string `json:"description"`

	Http2Enable bool `json:"http2_enable"`

	Id string `json:"id"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers"`

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	Protocol string `json:"protocol"`

	ProtocolPort int32 `json:"protocol_port"`

	SniContainerRefs []string `json:"sni_container_refs"`

	SniMatchAlgo string `json:"sni_match_algo"`

	Tags []Tag `json:"tags"`

	UpdatedAt string `json:"updated_at"`

	TlsCiphersPolicy string `json:"tls_ciphers_policy"`

	SecurityPolicyId string `json:"security_policy_id"`

	EnableMemberRetry bool `json:"enable_member_retry"`

	KeepaliveTimeout int32 `json:"keepalive_timeout"`

	ClientTimeout int32 `json:"client_timeout"`

	MemberTimeout int32 `json:"member_timeout"`

	Ipgroup *ListenerIpGroup `json:"ipgroup"`

	TransparentClientIpEnable bool `json:"transparent_client_ip_enable"`

	EnhanceL7policyEnable bool `json:"enhance_l7policy_enable"`

	QuicConfig *ListenerQuicConfig `json:"quic_config,omitempty"`

	ProtectionStatus *ListenerProtectionStatus `json:"protection_status,omitempty"`

	ProtectionReason *string `json:"protection_reason,omitempty"`

	GzipEnable *bool `json:"gzip_enable,omitempty"`

	PortRanges *[]PortRange `json:"port_ranges,omitempty"`
}

func (o Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}

type ListenerProtectionStatus struct {
	value string
}

type ListenerProtectionStatusEnum struct {
	NON_PROTECTION     ListenerProtectionStatus
	CONSOLE_PROTECTION ListenerProtectionStatus
}

func GetListenerProtectionStatusEnum() ListenerProtectionStatusEnum {
	return ListenerProtectionStatusEnum{
		NON_PROTECTION: ListenerProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: ListenerProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c ListenerProtectionStatus) Value() string {
	return c.value
}

func (c ListenerProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerProtectionStatus) UnmarshalJSON(b []byte) error {
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

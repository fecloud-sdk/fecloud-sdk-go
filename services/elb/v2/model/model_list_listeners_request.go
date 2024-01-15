package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenersRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	ConnectionLimit *int32 `json:"connection_limit,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	TlsContainerId *string `json:"tls_container_id,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}

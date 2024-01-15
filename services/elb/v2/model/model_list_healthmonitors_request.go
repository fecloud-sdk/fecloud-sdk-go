package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHealthmonitorsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Delay *int32 `json:"delay,omitempty"`

	MaxRetries *int32 `json:"max_retries,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Timeout *int32 `json:"timeout,omitempty"`

	Type *string `json:"type,omitempty"`

	MonitorPort *int32 `json:"monitor_port,omitempty"`

	ExpectedCodes *string `json:"expected_codes,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	UrlPath *string `json:"url_path,omitempty"`

	HttpMethod *string `json:"http_method,omitempty"`
}

func (o ListHealthmonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthmonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthmonitorsRequest", string(data)}, " ")
}

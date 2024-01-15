package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateHealthmonitorReq struct {
	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	MonitorPort *int32 `json:"monitor_port,omitempty"`

	Timeout *int32 `json:"timeout,omitempty"`

	ExpectedCodes *string `json:"expected_codes,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	UrlPath *string `json:"url_path,omitempty"`

	HttpMethod *string `json:"http_method,omitempty"`

	Delay *int32 `json:"delay,omitempty"`

	MaxRetries *int32 `json:"max_retries,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o UpdateHealthmonitorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorReq struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorReq", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchHostsProtectStatusRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SwitchHostsProtectStatusRequestInfo `json:"body,omitempty"`
}

func (o SwitchHostsProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHostsProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"SwitchHostsProtectStatusRequest", string(data)}, " ")
}

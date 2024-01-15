package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetWtpProtectionStatusInfoRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SetWtpProtectionStatusRequestInfo `json:"body,omitempty"`
}

func (o SetWtpProtectionStatusInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWtpProtectionStatusInfoRequest struct{}"
	}

	return strings.Join([]string{"SetWtpProtectionStatusInfoRequest", string(data)}, " ")
}

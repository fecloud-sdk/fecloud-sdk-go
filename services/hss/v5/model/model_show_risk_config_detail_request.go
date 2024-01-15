package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRiskConfigDetailRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CheckName string `json:"check_name"`

	Standard string `json:"standard"`

	HostId *string `json:"host_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowRiskConfigDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskConfigDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowRiskConfigDetailRequest", string(data)}, " ")
}

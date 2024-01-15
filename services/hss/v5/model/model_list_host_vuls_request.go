package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostVulsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostId string `json:"host_id"`

	Type *string `json:"type,omitempty"`

	VulName *string `json:"vul_name,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	HandleStatus *string `json:"handle_status,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o ListHostVulsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostVulsRequest struct{}"
	}

	return strings.Join([]string{"ListHostVulsRequest", string(data)}, " ")
}

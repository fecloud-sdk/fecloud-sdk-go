package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppsRequest struct {
	HostId string `json:"host_id"`

	HostName *string `json:"host_name,omitempty"`

	AppName *string `json:"app_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Version *string `json:"version,omitempty"`

	InstallDir *string `json:"install_dir,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}

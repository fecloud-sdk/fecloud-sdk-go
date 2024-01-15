package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListJarPackageHostInfoRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	FileName string `json:"file_name"`

	Category *string `json:"category,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListJarPackageHostInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageHostInfoRequest struct{}"
	}

	return strings.Join([]string{"ListJarPackageHostInfoRequest", string(data)}, " ")
}

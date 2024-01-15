package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListWeakPasswordUsersRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWeakPasswordUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWeakPasswordUsersRequest struct{}"
	}

	return strings.Join([]string{"ListWeakPasswordUsersRequest", string(data)}, " ")
}

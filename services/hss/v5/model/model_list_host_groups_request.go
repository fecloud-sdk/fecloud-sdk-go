package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostGroupsRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	GroupName *string `json:"group_name,omitempty"`
}

func (o ListHostGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListHostGroupsRequest", string(data)}, " ")
}

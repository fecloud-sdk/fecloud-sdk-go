package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAutoLaunchStatisticsRequest struct {
	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAutoLaunchStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchStatisticsRequest", string(data)}, " ")
}

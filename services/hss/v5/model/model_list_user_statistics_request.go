package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUserStatisticsRequest struct {
	UserName *string `json:"user_name,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUserStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListUserStatisticsRequest", string(data)}, " ")
}

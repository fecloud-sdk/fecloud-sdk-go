package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProcessStatisticsRequest struct {
	Path *string `json:"path,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProcessStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProcessStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListProcessStatisticsRequest", string(data)}, " ")
}

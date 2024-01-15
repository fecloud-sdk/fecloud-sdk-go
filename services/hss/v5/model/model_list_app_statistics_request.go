package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppStatisticsRequest struct {
	AppName *string `json:"app_name,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAppStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListAppStatisticsRequest", string(data)}, " ")
}

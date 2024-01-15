package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListJarPackageStatisticsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	Category *string `json:"category,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListJarPackageStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListJarPackageStatisticsRequest", string(data)}, " ")
}

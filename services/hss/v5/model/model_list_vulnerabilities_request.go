package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVulnerabilitiesRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Type *string `json:"type,omitempty"`

	VulId *string `json:"vul_id,omitempty"`

	VulName *string `json:"vul_name,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	RepairPriority *string `json:"repair_priority,omitempty"`

	HandleStatus *string `json:"handle_status,omitempty"`

	CveId *string `json:"cve_id,omitempty"`

	LabelList *string `json:"label_list,omitempty"`

	Status *string `json:"status,omitempty"`

	AssetValue *string `json:"asset_value,omitempty"`

	GroupName *string `json:"group_name,omitempty"`
}

func (o ListVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListVulnerabilitiesRequest", string(data)}, " ")
}

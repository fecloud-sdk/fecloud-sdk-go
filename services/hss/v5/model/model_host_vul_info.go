package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostVulInfo struct {
	VulName *string `json:"vul_name,omitempty"`

	VulId *string `json:"vul_id,omitempty"`

	LabelList *[]string `json:"label_list,omitempty"`

	RepairNecessity *string `json:"repair_necessity,omitempty"`

	ScanTime *int64 `json:"scan_time,omitempty"`

	Type *string `json:"type,omitempty"`

	AppList *[]HostVulInfoAppList `json:"app_list,omitempty"`

	SeverityLevel *string `json:"severity_level,omitempty"`

	SolutionDetail *string `json:"solution_detail,omitempty"`

	Url *string `json:"url,omitempty"`

	Description *string `json:"description,omitempty"`

	RepairCmd *string `json:"repair_cmd,omitempty"`

	Status *string `json:"status,omitempty"`

	RepairSuccessNum *int32 `json:"repair_success_num,omitempty"`

	CveList *[]HostVulInfoCveList `json:"cve_list,omitempty"`

	IsAffectBusiness *bool `json:"is_affect_business,omitempty"`

	FirstScanTime *int64 `json:"first_scan_time,omitempty"`
}

func (o HostVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulInfo struct{}"
	}

	return strings.Join([]string{"HostVulInfo", string(data)}, " ")
}

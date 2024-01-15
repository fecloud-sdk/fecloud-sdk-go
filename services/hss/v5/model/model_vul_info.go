package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VulInfo struct {
	VulName *string `json:"vul_name,omitempty"`

	VulId *string `json:"vul_id,omitempty"`

	LabelList *[]string `json:"label_list,omitempty"`

	RepairNecessity *string `json:"repair_necessity,omitempty"`

	SeverityLevel *string `json:"severity_level,omitempty"`

	HostNum *int32 `json:"host_num,omitempty"`

	UnhandleHostNum *int32 `json:"unhandle_host_num,omitempty"`

	ScanTime *int64 `json:"scan_time,omitempty"`

	SolutionDetail *string `json:"solution_detail,omitempty"`

	Url *string `json:"url,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *string `json:"type,omitempty"`

	HostIdList *[]string `json:"host_id_list,omitempty"`

	CveList *[]VulInfoCveList `json:"cve_list,omitempty"`

	PatchUrl *string `json:"patch_url,omitempty"`

	RepairPriority *string `json:"repair_priority,omitempty"`

	HostsNum *VulnerabilityHostNumberInfo `json:"hosts_num,omitempty"`

	RepairSuccessNum *int32 `json:"repair_success_num,omitempty"`

	FixedNum *int64 `json:"fixed_num,omitempty"`

	IgnoredNum *int64 `json:"ignored_num,omitempty"`

	VerifyNum *int32 `json:"verify_num,omitempty"`
}

func (o VulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulInfo struct{}"
	}

	return strings.Join([]string{"VulInfo", string(data)}, " ")
}

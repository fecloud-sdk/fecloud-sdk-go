package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostVulInfoCveList struct {
	CveId *string `json:"cve_id,omitempty"`

	Cvss *float32 `json:"cvss,omitempty"`
}

func (o HostVulInfoCveList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulInfoCveList struct{}"
	}

	return strings.Join([]string{"HostVulInfoCveList", string(data)}, " ")
}

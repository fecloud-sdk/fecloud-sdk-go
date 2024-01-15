package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VulInfoCveList struct {
	CveId *string `json:"cve_id,omitempty"`

	Cvss *float32 `json:"cvss,omitempty"`
}

func (o VulInfoCveList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulInfoCveList struct{}"
	}

	return strings.Join([]string{"VulInfoCveList", string(data)}, " ")
}

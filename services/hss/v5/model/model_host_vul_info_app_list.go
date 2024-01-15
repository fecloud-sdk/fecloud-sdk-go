package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostVulInfoAppList struct {
	AppName *string `json:"app_name,omitempty"`

	AppVersion *string `json:"app_version,omitempty"`

	UpgradeVersion *string `json:"upgrade_version,omitempty"`

	AppPath *string `json:"app_path,omitempty"`
}

func (o HostVulInfoAppList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulInfoAppList struct{}"
	}

	return strings.Join([]string{"HostVulInfoAppList", string(data)}, " ")
}

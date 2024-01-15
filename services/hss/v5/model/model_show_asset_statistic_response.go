package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowAssetStatisticResponse struct {
	AccountNum *int64 `json:"account_num,omitempty"`

	PortNum *int64 `json:"port_num,omitempty"`

	ProcessNum *int64 `json:"process_num,omitempty"`

	AppNum *int64 `json:"app_num,omitempty"`

	AutoLaunchNum *int64 `json:"auto_launch_num,omitempty"`

	WebFrameworkNum *int64 `json:"web_framework_num,omitempty"`

	WebSiteNum *int64 `json:"web_site_num,omitempty"`

	JarPackageNum *int64 `json:"jar_package_num,omitempty"`

	KernelModuleNum *int64 `json:"kernel_module_num,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowAssetStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetStatisticResponse", string(data)}, " ")
}

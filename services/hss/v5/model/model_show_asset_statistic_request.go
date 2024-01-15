package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowAssetStatisticRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HostId *string `json:"host_id,omitempty"`
}

func (o ShowAssetStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetStatisticRequest", string(data)}, " ")
}

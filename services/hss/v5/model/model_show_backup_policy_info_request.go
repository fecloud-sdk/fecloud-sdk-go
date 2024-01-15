package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupPolicyInfoRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowBackupPolicyInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyInfoRequest", string(data)}, " ")
}

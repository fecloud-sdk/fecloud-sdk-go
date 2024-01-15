package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateBackupPolicyInfoRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateBackupPolicyRequestInfo `json:"body,omitempty"`
}

func (o UpdateBackupPolicyInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyInfoRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupPolicyResponse struct {
	BackupPolicy   *ShowBackupPolicyResult `json:"backup_policy,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyResponse", string(data)}, " ")
}

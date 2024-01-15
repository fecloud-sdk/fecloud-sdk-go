package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetBackupPolicyRequestBody struct {
	BackupPolicy *BackupPolicy `json:"backup_policy"`

	ReserveBackups *bool `json:"reserve_backups,omitempty"`
}

func (o SetBackupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetBackupPolicyRequestBody", string(data)}, " ")
}

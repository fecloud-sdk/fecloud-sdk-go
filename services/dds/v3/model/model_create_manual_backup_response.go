package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateManualBackupResponse struct {
	JobId *string `json:"job_id,omitempty"`

	BackupId       *string `json:"backup_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}

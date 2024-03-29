package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateManualBackupResponse struct {
	Backup         *BackupInfo `json:"backup,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}

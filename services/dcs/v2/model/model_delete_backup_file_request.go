package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteBackupFileRequest struct {
	BackupId string `json:"backup_id"`

	InstanceId string `json:"instance_id"`
}

func (o DeleteBackupFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupFileRequest", string(data)}, " ")
}

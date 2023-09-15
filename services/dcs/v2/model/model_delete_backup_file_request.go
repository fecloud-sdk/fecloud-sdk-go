package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteBackupFileRequest Request Object
type DeleteBackupFileRequest struct {

	// 备份记录ID。
	BackupId string `json:"backup_id"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeleteBackupFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupFileRequest", string(data)}, " ")
}

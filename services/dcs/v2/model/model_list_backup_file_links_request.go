package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBackupFileLinksRequest struct {
	InstanceId string `json:"instance_id"`

	BackupId string `json:"backup_id"`

	Body *DownloadBackupFilesReq `json:"body,omitempty"`
}

func (o ListBackupFileLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupFileLinksRequest struct{}"
	}

	return strings.Join([]string{"ListBackupFileLinksRequest", string(data)}, " ")
}

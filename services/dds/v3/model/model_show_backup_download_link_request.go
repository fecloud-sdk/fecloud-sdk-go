package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupDownloadLinkRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	BackupId string `json:"backup_id"`
}

func (o ShowBackupDownloadLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupDownloadLinkRequest", string(data)}, " ")
}

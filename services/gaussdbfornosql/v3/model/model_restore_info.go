package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreInfo struct {
	BackupId *string `json:"backup_id,omitempty"`

	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	RestoreTime *int64 `json:"restore_time,omitempty"`
}

func (o RestoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInfo struct{}"
	}

	return strings.Join([]string{"RestoreInfo", string(data)}, " ")
}

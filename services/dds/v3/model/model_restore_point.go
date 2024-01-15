package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestorePoint struct {
	InstanceId *string `json:"instance_id,omitempty"`

	Type *string `json:"type,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	RestoreTime *string `json:"restore_time,omitempty"`
}

func (o RestorePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePoint struct{}"
	}

	return strings.Join([]string{"RestorePoint", string(data)}, " ")
}

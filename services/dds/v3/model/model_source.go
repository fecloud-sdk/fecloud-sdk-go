package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Source struct {
	InstanceId string `json:"instance_id"`

	Type *string `json:"type,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	RestoreTime *string `json:"restore_time,omitempty"`
}

func (o Source) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Source struct{}"
	}

	return strings.Join([]string{"Source", string(data)}, " ")
}

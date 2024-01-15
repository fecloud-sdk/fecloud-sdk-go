package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupTriggerPropertiesInfo struct {
	Pattern *[]string `json:"pattern,omitempty"`

	StartTime *string `json:"start_time,omitempty"`
}

func (o BackupTriggerPropertiesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerPropertiesInfo struct{}"
	}

	return strings.Join([]string{"BackupTriggerPropertiesInfo", string(data)}, " ")
}

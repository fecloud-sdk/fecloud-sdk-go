package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupTriggerInfo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`

	Properties *BackupTriggerPropertiesInfo `json:"properties,omitempty"`
}

func (o BackupTriggerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerInfo struct{}"
	}

	return strings.Join([]string{"BackupTriggerInfo", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupTriggerRequestInfo1 struct {
	Properties *BackupTriggerPropertiesRequestInfo1 `json:"properties,omitempty"`
}

func (o BackupTriggerRequestInfo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerRequestInfo1 struct{}"
	}

	return strings.Join([]string{"BackupTriggerRequestInfo1", string(data)}, " ")
}

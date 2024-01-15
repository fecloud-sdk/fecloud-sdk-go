package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupTriggerRequestInfo struct {
	Properties *BackupTriggerPropertiesRequestInfo `json:"properties"`
}

func (o BackupTriggerRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerRequestInfo struct{}"
	}

	return strings.Join([]string{"BackupTriggerRequestInfo", string(data)}, " ")
}

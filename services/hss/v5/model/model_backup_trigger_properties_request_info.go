package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupTriggerPropertiesRequestInfo struct {
	Pattern []string `json:"pattern"`
}

func (o BackupTriggerPropertiesRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerPropertiesRequestInfo struct{}"
	}

	return strings.Join([]string{"BackupTriggerPropertiesRequestInfo", string(data)}, " ")
}

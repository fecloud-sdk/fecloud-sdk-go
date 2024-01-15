package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupTriggerPropertiesRequestInfo1 struct {
	Pattern *[]string `json:"pattern,omitempty"`
}

func (o BackupTriggerPropertiesRequestInfo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerPropertiesRequestInfo1 struct{}"
	}

	return strings.Join([]string{"BackupTriggerPropertiesRequestInfo1", string(data)}, " ")
}

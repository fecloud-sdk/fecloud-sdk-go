package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateBackupPolicyRequestInfo struct {
	Enabled *bool `json:"enabled,omitempty"`

	PolicyId string `json:"policy_id"`

	OperationDefinition *OperationDefinitionRequestInfo `json:"operation_definition,omitempty"`

	Trigger *BackupTriggerRequestInfo `json:"trigger,omitempty"`
}

func (o UpdateBackupPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyRequestInfo", string(data)}, " ")
}

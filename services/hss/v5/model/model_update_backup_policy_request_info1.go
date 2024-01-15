package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateBackupPolicyRequestInfo1 struct {
	Enabled *bool `json:"enabled,omitempty"`

	PolicyId *string `json:"policy_id,omitempty"`

	OperationDefinition *OperationDefinitionRequestInfo `json:"operation_definition,omitempty"`

	Trigger *BackupTriggerRequestInfo1 `json:"trigger,omitempty"`
}

func (o UpdateBackupPolicyRequestInfo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupPolicyRequestInfo1 struct{}"
	}

	return strings.Join([]string{"UpdateBackupPolicyRequestInfo1", string(data)}, " ")
}

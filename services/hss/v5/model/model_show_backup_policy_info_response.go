package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupPolicyInfoResponse struct {
	Enabled *bool `json:"enabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	OperationType *string `json:"operation_type,omitempty"`

	OperationDefinition *OperationDefinitionInfo `json:"operation_definition,omitempty"`

	Trigger        *BackupTriggerInfo `json:"trigger,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowBackupPolicyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyInfoResponse", string(data)}, " ")
}

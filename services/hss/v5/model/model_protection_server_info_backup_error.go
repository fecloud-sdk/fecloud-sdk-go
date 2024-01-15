package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProtectionServerInfoBackupError struct {
	ErrorCode *int32 `json:"error_code,omitempty"`

	ErrorDescription *string `json:"error_description,omitempty"`
}

func (o ProtectionServerInfoBackupError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionServerInfoBackupError struct{}"
	}

	return strings.Join([]string{"ProtectionServerInfoBackupError", string(data)}, " ")
}

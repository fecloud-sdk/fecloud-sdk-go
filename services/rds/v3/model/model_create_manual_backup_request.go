package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateManualBackupRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateManualBackupRequestBody `json:"body,omitempty"`
}

func (o CreateManualBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateManualBackupRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateManualBackupRequestBody struct {
	Backup *CreateManualBackupOption `json:"backup"`
}

func (o CreateManualBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateManualBackupRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateManualBackupOption struct {
	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`
}

func (o CreateManualBackupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupOption struct{}"
	}

	return strings.Join([]string{"CreateManualBackupOption", string(data)}, " ")
}

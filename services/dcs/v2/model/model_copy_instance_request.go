package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CopyInstanceRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BackupInstanceBody `json:"body,omitempty"`
}

func (o CopyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceRequest struct{}"
	}

	return strings.Join([]string{"CopyInstanceRequest", string(data)}, " ")
}

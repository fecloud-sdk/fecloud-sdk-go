package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BackupResources struct {
	VaultId *string `json:"vault_id,omitempty"`

	ResourceList *[]ResourceInfo `json:"resource_list,omitempty"`
}

func (o BackupResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupResources struct{}"
	}

	return strings.Join([]string{"BackupResources", string(data)}, " ")
}

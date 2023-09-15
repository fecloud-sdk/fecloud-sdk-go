package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CopyInstanceResponse Response Object
type CopyInstanceResponse struct {

	// 备份记录ID。
	BackupId       *string `json:"backup_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceResponse struct{}"
	}

	return strings.Join([]string{"CopyInstanceResponse", string(data)}, " ")
}

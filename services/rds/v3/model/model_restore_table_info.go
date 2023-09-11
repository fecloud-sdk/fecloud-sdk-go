package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreTableInfo struct {

	// 旧表名
	OldName string `json:"oldName"`

	// 新表名
	NewName string `json:"newName"`
}

func (o RestoreTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTableInfo struct{}"
	}

	return strings.Join([]string{"RestoreTableInfo", string(data)}, " ")
}

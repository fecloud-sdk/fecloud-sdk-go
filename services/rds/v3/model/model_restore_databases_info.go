package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreDatabasesInfo struct {

	// 库名
	Database string `json:"database"`

	// 表信息
	Tables []RestoreTableInfo `json:"tables"`
}

func (o RestoreDatabasesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDatabasesInfo struct{}"
	}

	return strings.Join([]string{"RestoreDatabasesInfo", string(data)}, " ")
}

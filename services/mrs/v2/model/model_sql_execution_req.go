package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SqlExecutionReq struct {
	SqlType string `json:"sql_type"`

	SqlContent string `json:"sql_content"`

	Database *string `json:"database,omitempty"`

	ArchivePath *string `json:"archive_path,omitempty"`
}

func (o SqlExecutionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlExecutionReq struct{}"
	}

	return strings.Join([]string{"SqlExecutionReq", string(data)}, " ")
}

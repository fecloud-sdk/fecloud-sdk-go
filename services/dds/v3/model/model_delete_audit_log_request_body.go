package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAuditLogRequestBody struct {

	// 文件名列表
	FileNames []string `json:"file_names"`
}

func (o DeleteAuditLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditLogRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteAuditLogRequestBody", string(data)}, " ")
}

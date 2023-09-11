package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BinlogClearPolicyRequestBody struct {

	// 取值范围0-7*24
	BinlogRetentionHours int64 `json:"binlog_retention_hours"`
}

func (o BinlogClearPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinlogClearPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"BinlogClearPolicyRequestBody", string(data)}, " ")
}

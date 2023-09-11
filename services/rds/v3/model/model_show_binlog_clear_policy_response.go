package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowBinlogClearPolicyResponse Response Object
type ShowBinlogClearPolicyResponse struct {

	// binlog保留时长
	BinlogRetentionHours *int32 `json:"binlog_retention_hours,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowBinlogClearPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogClearPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowBinlogClearPolicyResponse", string(data)}, " ")
}

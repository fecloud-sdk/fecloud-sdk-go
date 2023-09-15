package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListBackupRecordsResponse Response Object
type ListBackupRecordsResponse struct {

	// 返回记录数。
	TotalNum *int32 `json:"total_num,omitempty"`

	// 备份信息的详情数组。
	BackupRecordResponse *[]BackupRecordResponse `json:"backup_record_response,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListBackupRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupRecordsResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteMigrateTaskRequest 删除迁移任务请求体
type DeleteMigrateTaskRequest struct {

	// 删除的迁移任务ID列表。
	TaskIdList []string `json:"task_id_list"`
}

func (o DeleteMigrateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrateTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrateTaskRequest", string(data)}, " ")
}

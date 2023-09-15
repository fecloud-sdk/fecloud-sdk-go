package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchStopMigrationTasksResponse Response Object
type BatchStopMigrationTasksResponse struct {

	// 数据迁移任务列表。
	MigrationTasks *[]StopMigrationTaskResult `json:"migration_tasks,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchStopMigrationTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopMigrationTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchStopMigrationTasksResponse", string(data)}, " ")
}

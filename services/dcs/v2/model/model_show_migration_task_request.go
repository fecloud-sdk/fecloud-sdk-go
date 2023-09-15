package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowMigrationTaskRequest Request Object
type ShowMigrationTaskRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskRequest", string(data)}, " ")
}

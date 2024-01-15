package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMigrationTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskRequest", string(data)}, " ")
}

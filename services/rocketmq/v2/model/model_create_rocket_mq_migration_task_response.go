package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRocketMqMigrationTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskResponse", string(data)}, " ")
}

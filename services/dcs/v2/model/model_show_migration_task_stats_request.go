package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMigrationTaskStatsRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskStatsRequest", string(data)}, " ")
}

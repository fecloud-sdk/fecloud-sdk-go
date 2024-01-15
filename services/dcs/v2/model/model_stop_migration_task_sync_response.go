package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StopMigrationTaskSyncResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopMigrationTaskSyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskSyncResponse struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskSyncResponse", string(data)}, " ")
}

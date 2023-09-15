package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteMigrationTaskRequest Request Object
type DeleteMigrationTaskRequest struct {
	Body *DeleteMigrateTaskRequest `json:"body,omitempty"`
}

func (o DeleteMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrationTaskRequest", string(data)}, " ")
}

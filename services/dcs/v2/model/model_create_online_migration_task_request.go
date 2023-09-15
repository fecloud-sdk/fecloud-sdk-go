package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateOnlineMigrationTaskRequest Request Object
type CreateOnlineMigrationTaskRequest struct {
	Body *CreateOnlineMigrationTaskBody `json:"body,omitempty"`
}

func (o CreateOnlineMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnlineMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateOnlineMigrationTaskRequest", string(data)}, " ")
}

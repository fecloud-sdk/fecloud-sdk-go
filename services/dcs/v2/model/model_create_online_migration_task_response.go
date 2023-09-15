package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateOnlineMigrationTaskResponse Response Object
type CreateOnlineMigrationTaskResponse struct {

	// 在线迁移任务ID。
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOnlineMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnlineMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateOnlineMigrationTaskResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateDatabaseRequest Request Object
type CreateDatabaseRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DatabaseForCreation `json:"body,omitempty"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}

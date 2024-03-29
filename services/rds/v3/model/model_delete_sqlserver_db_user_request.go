package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSqlserverDbUserRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`
}

func (o DeleteSqlserverDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDbUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDbUserRequest", string(data)}, " ")
}

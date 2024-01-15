package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteRocketMqMigrationTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteRocketMqMigrationTaskResponse", string(data)}, " ")
}

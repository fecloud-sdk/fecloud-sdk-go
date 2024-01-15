package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRocketMqMigrationTaskResponse struct {
	Total *int32 `json:"total,omitempty"`

	Task           *[]MetadataTask `json:"task,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ListRocketMqMigrationTaskResponse", string(data)}, " ")
}

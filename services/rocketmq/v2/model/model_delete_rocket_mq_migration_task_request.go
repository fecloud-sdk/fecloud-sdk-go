package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteRocketMqMigrationTaskRequest struct {
	InstanceId string `json:"instance_id"`

	Body *MetadataDeleteReq `json:"body,omitempty"`
}

func (o DeleteRocketMqMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRocketMqMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteRocketMqMigrationTaskRequest", string(data)}, " ")
}

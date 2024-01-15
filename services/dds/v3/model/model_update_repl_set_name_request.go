package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateReplSetNameRequest struct {
	InstanceId string `json:"instance_id"`

	Body *MongoUpdateReplSetV3RequestBody `json:"body,omitempty"`
}

func (o UpdateReplSetNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplSetNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateReplSetNameRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdatePasswordRequest Request Object
type UpdatePasswordRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyInstancePasswordBody `json:"body,omitempty"`
}

func (o UpdatePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdatePasswordRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateIpWhitelistRequest Request Object
type UpdateIpWhitelistRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyIpWhitelistBody `json:"body,omitempty"`
}

func (o UpdateIpWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpWhitelistRequest", string(data)}, " ")
}

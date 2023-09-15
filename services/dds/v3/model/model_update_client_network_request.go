package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateClientNetworkRequest Request Object
type UpdateClientNetworkRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ClientNetworkRequestBody `json:"body,omitempty"`
}

func (o UpdateClientNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkRequest struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateBigkeyAutoscanConfigRequest Request Object
type UpdateBigkeyAutoscanConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *AutoscanConfigRequest `json:"body,omitempty"`
}

func (o UpdateBigkeyAutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBigkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateBigkeyAutoscanConfigRequest", string(data)}, " ")
}

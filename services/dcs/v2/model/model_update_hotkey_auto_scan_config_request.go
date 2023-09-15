package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateHotkeyAutoScanConfigRequest Request Object
type UpdateHotkeyAutoScanConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *AutoscanConfigRequest `json:"body,omitempty"`
}

func (o UpdateHotkeyAutoScanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotkeyAutoScanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateHotkeyAutoScanConfigRequest", string(data)}, " ")
}

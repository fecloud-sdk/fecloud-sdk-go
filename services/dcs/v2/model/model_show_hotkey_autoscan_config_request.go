package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowHotkeyAutoscanConfigRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowHotkeyAutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowHotkeyAutoscanConfigRequest", string(data)}, " ")
}

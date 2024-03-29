package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetReadOnlySwitchRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *MysqlReadOnlySwitch `json:"body,omitempty"`
}

func (o SetReadOnlySwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetReadOnlySwitchRequest struct{}"
	}

	return strings.Join([]string{"SetReadOnlySwitchRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchEnableAlarmRulesRequest struct {
	ContentType string `json:"Content-Type"`

	Body *BatchEnableAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchEnableAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmRulesRequest", string(data)}, " ")
}

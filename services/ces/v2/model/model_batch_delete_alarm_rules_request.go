package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteAlarmRulesRequest struct {
	ContentType string `json:"Content-Type"`

	Body *BatchDeleteAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmRulesRequest", string(data)}, " ")
}

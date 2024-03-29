package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAlarmRuleResourcesRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId string `json:"alarm_id"`

	Body *ResourcesReqV2 `json:"body,omitempty"`
}

func (o DeleteAlarmRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleResourcesRequest", string(data)}, " ")
}

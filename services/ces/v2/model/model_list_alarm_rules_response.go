package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmRulesResponse struct {
	Alarms *[]ListAlarmResponseAlarms `json:"alarms,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesResponse", string(data)}, " ")
}

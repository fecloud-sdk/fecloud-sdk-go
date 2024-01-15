package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId string `json:"alarm_id"`

	Body *UpdateAlarmRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRequest", string(data)}, " ")
}

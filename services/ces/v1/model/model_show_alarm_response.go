package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowAlarmResponse struct {
	MetricAlarms   *[]MetricAlarms `json:"metric_alarms,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmResponse", string(data)}, " ")
}

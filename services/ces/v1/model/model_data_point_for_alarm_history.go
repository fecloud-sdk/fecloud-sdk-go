package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DataPointForAlarmHistory struct {
	Time *int64 `json:"time,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func (o DataPointForAlarmHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointForAlarmHistory struct{}"
	}

	return strings.Join([]string{"DataPointForAlarmHistory", string(data)}, " ")
}

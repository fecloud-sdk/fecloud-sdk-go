package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OperationDefinitionRequestInfo struct {
	DayBackups *int32 `json:"day_backups,omitempty"`

	MaxBackups *int32 `json:"max_backups,omitempty"`

	MonthBackups *int32 `json:"month_backups,omitempty"`

	RetentionDurationDays *int32 `json:"retention_duration_days,omitempty"`

	Timezone *string `json:"timezone,omitempty"`

	WeekBackups *int32 `json:"week_backups,omitempty"`

	YearBackups *int32 `json:"year_backups,omitempty"`
}

func (o OperationDefinitionRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationDefinitionRequestInfo struct{}"
	}

	return strings.Join([]string{"OperationDefinitionRequestInfo", string(data)}, " ")
}

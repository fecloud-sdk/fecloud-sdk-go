package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RecyclePolicy struct {
	Enabled bool `json:"enabled"`

	RetentionPeriodInDays *int32 `json:"retention_period_in_days,omitempty"`
}

func (o RecyclePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicy struct{}"
	}

	return strings.Join([]string{"RecyclePolicy", string(data)}, " ")
}

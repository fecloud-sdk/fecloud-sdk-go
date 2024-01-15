package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourcesPlan struct {
	PeriodType string `json:"period_type"`

	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	MinCapacity int32 `json:"min_capacity"`

	MaxCapacity int32 `json:"max_capacity"`
}

func (o ResourcesPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesPlan struct{}"
	}

	return strings.Join([]string{"ResourcesPlan", string(data)}, " ")
}

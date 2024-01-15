package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Trigger struct {
	MetricName string `json:"metric_name"`

	MetricValue string `json:"metric_value"`

	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	EvaluationPeriods int32 `json:"evaluation_periods"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}

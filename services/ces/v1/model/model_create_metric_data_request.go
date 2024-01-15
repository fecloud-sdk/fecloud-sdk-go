package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateMetricDataRequest struct {
	Body *[]MetricDataItem `json:"body,omitempty"`
}

func (o CreateMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetricDataRequest struct{}"
	}

	return strings.Join([]string{"CreateMetricDataRequest", string(data)}, " ")
}

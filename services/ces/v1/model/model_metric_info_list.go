package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MetricInfoList struct {
	Dimensions []MetricsDimension `json:"dimensions"`

	MetricName string `json:"metric_name"`

	Namespace string `json:"namespace"`

	Unit string `json:"unit"`
}

func (o MetricInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfoList struct{}"
	}

	return strings.Join([]string{"MetricInfoList", string(data)}, " ")
}

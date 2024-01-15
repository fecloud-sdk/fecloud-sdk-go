package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MetricName struct {
}

func (o MetricName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricName struct{}"
	}

	return strings.Join([]string{"MetricName", string(data)}, " ")
}

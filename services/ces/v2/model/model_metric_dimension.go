package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MetricDimension struct {
	Name string `json:"name"`

	Value *string `json:"value,omitempty"`
}

func (o MetricDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimension struct{}"
	}

	return strings.Join([]string{"MetricDimension", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DiffDetails struct {
	ParameterName string `json:"parameter_name"`

	SourceValue string `json:"source_value"`

	TargetValue string `json:"target_value"`
}

func (o DiffDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffDetails struct{}"
	}

	return strings.Join([]string{"DiffDetails", string(data)}, " ")
}

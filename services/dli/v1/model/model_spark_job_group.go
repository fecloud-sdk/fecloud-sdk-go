package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SparkJobGroup struct {
	Name *string `json:"name,omitempty"`

	Resources *[]SparkJobResource `json:"resources,omitempty"`
}

func (o SparkJobGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobGroup struct{}"
	}

	return strings.Join([]string{"SparkJobGroup", string(data)}, " ")
}

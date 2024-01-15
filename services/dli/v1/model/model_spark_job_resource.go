package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SparkJobResource struct {
	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o SparkJobResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobResource struct{}"
	}

	return strings.Join([]string{"SparkJobResource", string(data)}, " ")
}

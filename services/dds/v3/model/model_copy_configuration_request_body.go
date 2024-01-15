package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CopyConfigurationRequestBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`
}

func (o CopyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequestBody", string(data)}, " ")
}

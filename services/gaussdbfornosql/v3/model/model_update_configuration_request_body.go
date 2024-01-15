package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationRequestBody struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Values map[string]string `json:"values,omitempty"`
}

func (o UpdateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequestBody", string(data)}, " ")
}

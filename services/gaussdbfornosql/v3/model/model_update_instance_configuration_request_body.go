package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceConfigurationRequestBody struct {
	Values map[string]string `json:"values"`
}

func (o UpdateInstanceConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationRequestBody struct {
	InstanceIds []string `json:"instance_ids"`
}

func (o ApplyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequestBody", string(data)}, " ")
}

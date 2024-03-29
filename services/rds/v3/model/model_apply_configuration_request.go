package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationRequest struct {
	InstanceIds []string `json:"instance_ids"`
}

func (o ApplyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationRspConfiguration struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	IgnoredParams *[]string `json:"ignored_params,omitempty"`
}

func (o UpdateConfigurationRspConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRspConfiguration struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRspConfiguration", string(data)}, " ")
}

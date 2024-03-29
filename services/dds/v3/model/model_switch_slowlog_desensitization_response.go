package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchSlowlogDesensitizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSlowlogDesensitizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationResponse struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationResponse", string(data)}, " ")
}

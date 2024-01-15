package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSlowlogDesensitizationSwitchResponse struct {
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSlowlogDesensitizationSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowlogDesensitizationSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowlogDesensitizationSwitchResponse", string(data)}, " ")
}

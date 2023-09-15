package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateConfigurationParameterResponse Response Object
type UpdateConfigurationParameterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterResponse", string(data)}, " ")
}

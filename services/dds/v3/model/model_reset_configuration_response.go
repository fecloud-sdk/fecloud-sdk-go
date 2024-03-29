package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetConfigurationResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ResetConfigurationResponse", string(data)}, " ")
}

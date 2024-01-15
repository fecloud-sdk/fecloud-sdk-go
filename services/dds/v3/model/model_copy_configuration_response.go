package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CopyConfigurationResponse struct {
	ConfigurationId *string `json:"configuration_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CopyConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CopyConfigurationResponse", string(data)}, " ")
}

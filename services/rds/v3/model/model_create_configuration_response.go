package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateConfigurationResponse Response Object
type CreateConfigurationResponse struct {
	Configuration  *ConfigurationSummaryForCreate `json:"configuration,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CreateConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigurationResponse", string(data)}, " ")
}

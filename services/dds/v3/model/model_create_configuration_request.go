package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateConfigurationRequest Request Object
type CreateConfigurationRequest struct {
	Body *CreateConfigurationRequestBody `json:"body,omitempty"`
}

func (o CreateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequest", string(data)}, " ")
}

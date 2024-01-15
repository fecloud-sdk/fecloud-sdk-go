package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchConfigurationResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationResponse", string(data)}, " ")
}

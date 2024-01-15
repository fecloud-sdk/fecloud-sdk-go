package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEntityConfigurationResponse struct {
	JobId *string `json:"job_id,omitempty"`

	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateEntityConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEntityConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateEntityConfigurationResponse", string(data)}, " ")
}

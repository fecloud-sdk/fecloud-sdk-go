package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceConfigurationResponse struct {
	JobId *string `json:"job_id,omitempty"`

	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationResponse", string(data)}, " ")
}

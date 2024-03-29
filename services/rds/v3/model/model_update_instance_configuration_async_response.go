package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceConfigurationAsyncResponse struct {
	JobId *string `json:"job_id,omitempty"`

	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateInstanceConfigurationAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationAsyncResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationAsyncResponse", string(data)}, " ")
}

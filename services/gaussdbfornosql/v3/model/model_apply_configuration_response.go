package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApplyConfigurationResponse struct {
	JobId *string `json:"job_id,omitempty"`

	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ApplyConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationResponse", string(data)}, " ")
}

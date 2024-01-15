package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowFlinkJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobDetail      *FlinkJobDetail `json:"job_detail,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobResponse", string(data)}, " ")
}

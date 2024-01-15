package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSingleJobExeResponse struct {
	JobDetail      *JobQueryBean `json:"job_detail,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSingleJobExeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleJobExeResponse struct{}"
	}

	return strings.Join([]string{"ShowSingleJobExeResponse", string(data)}, " ")
}

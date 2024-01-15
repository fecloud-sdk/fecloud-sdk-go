package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStopServersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchStopServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServersResponse struct{}"
	}

	return strings.Join([]string{"BatchStopServersResponse", string(data)}, " ")
}

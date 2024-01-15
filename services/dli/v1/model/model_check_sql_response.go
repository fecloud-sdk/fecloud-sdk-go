package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckSqlResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobType        *string `json:"job_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSqlResponse struct{}"
	}

	return strings.Join([]string{"CheckSqlResponse", string(data)}, " ")
}

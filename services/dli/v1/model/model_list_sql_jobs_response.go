package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSqlJobsResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobCount *int32 `json:"job_count,omitempty"`

	Jobs           *[]SqlJobInfo `json:"jobs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListSqlJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobsResponse struct{}"
	}

	return strings.Join([]string{"ListSqlJobsResponse", string(data)}, " ")
}

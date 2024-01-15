package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlJobProgressResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	Status *string `json:"status,omitempty"`

	SubJobId *int32 `json:"sub_job_id,omitempty"`

	Progress *float64 `json:"progress,omitempty"`

	SubJobs        *[]SubJobInfo `json:"sub_jobs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSqlJobProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobProgressResponse", string(data)}, " ")
}

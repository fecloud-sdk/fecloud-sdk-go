package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlinkJobsResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobList        *FlinkJobInfoList `json:"job_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsResponse", string(data)}, " ")
}

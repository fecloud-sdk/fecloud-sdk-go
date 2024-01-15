package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobEntities struct {
	SubJobs *[]SubJob `json:"sub_jobs,omitempty"`

	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}

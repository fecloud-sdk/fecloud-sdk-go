package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlinkJobsRequest struct {
	JobType *string `json:"job_type,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int64 `json:"offset,omitempty"`

	Order *string `json:"order,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	RootJobId *int64 `json:"root_job_id,omitempty"`

	ShowDetail *bool `json:"show_detail,omitempty"`

	Status *string `json:"status,omitempty"`

	SysEnterpriseProjectName *string `json:"sys_enterprise_project_name,omitempty"`

	Tags *string `json:"tags,omitempty"`

	UserName *string `json:"user_name,omitempty"`
}

func (o ListFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsRequest", string(data)}, " ")
}

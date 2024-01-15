package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkJobDetail struct {
	JobId int64 `json:"job_id"`

	Name *string `json:"name,omitempty"`

	Desc *string `json:"desc,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusDesc *string `json:"status_desc,omitempty"`

	CreateTime int64 `json:"create_time"`

	StartTime *int64 `json:"start_time,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	SqlBody *string `json:"sql_body,omitempty"`

	RunMode *string `json:"run_mode,omitempty"`

	JobConfig *FlinkJobDetailConfig `json:"job_config,omitempty"`

	MainClass *string `json:"main_class,omitempty"`

	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	ExecutionGraph *string `json:"execution_graph,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	SavepointPath *string `json:"savepoint_path,omitempty"`
}

func (o FlinkJobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobDetail struct{}"
	}

	return strings.Join([]string{"FlinkJobDetail", string(data)}, " ")
}

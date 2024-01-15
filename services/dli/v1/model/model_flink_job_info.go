package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkJobInfo struct {
	JobId int64 `json:"job_id"`

	Name *string `json:"name,omitempty"`

	Desc *string `json:"desc,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusDesc *string `json:"status_desc,omitempty"`

	CreateTime int64 `json:"create_time"`

	StartTime *int64 `json:"start_time,omitempty"`

	Duration *int64 `json:"duration,omitempty"`

	RootId *int64 `json:"root_id,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	SqlBody *string `json:"sql_body,omitempty"`

	RunMode *string `json:"run_mode,omitempty"`

	JobConfig *FlinkJobConfigInfo `json:"job_config,omitempty"`

	MainClass *string `json:"main_class,omitempty"`

	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	ExecutionGraph *string `json:"execution_graph,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	GraphEditorEnabled *bool `json:"graph_editor_enabled,omitempty"`

	HasSavepoint *bool `json:"has_savepoint,omitempty"`
}

func (o FlinkJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobInfo struct{}"
	}

	return strings.Join([]string{"FlinkJobInfo", string(data)}, " ")
}

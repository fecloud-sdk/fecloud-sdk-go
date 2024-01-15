package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddJobsReqV11 struct {
	JobType int32 `json:"job_type"`

	JobName string `json:"job_name"`

	JarPath *string `json:"jar_path,omitempty"`

	Arguments *string `json:"arguments,omitempty"`

	Input *string `json:"input,omitempty"`

	Output *string `json:"output,omitempty"`

	JobLog *string `json:"job_log,omitempty"`

	HiveScriptPath *string `json:"hive_script_path,omitempty"`

	Hql *string `json:"hql,omitempty"`

	ShutdownCluster *bool `json:"shutdown_cluster,omitempty"`

	SubmitJobOnceClusterRun bool `json:"submit_job_once_cluster_run"`

	FileAction *string `json:"file_action,omitempty"`
}

func (o AddJobsReqV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddJobsReqV11 struct{}"
	}

	return strings.Join([]string{"AddJobsReqV11", string(data)}, " ")
}

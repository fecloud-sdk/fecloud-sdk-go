package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkJobConfigInfo struct {
	CheckpointEnabled *bool `json:"checkpoint_enabled,omitempty"`

	CheckpointMode *string `json:"checkpoint_mode,omitempty"`

	CheckpointInterval *int32 `json:"checkpoint_interval,omitempty"`

	LogEnabled *bool `json:"log_enabled,omitempty"`

	ObsBucket *string `json:"obs_bucket,omitempty"`

	SmnTopic *string `json:"smn_topic,omitempty"`

	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	RootId *int32 `json:"root_id,omitempty"`

	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	CuNumber *int32 `json:"cu_number,omitempty"`

	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	IdleStateRetention *int32 `json:"idle_state_retention,omitempty"`

	UdfJarUrl *string `json:"udf_jar_url,omitempty"`

	DirtyDataStrategy *string `json:"dirty_data_strategy,omitempty"`

	Entrypoint *string `json:"entrypoint,omitempty"`

	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	ExecutorNumber *int32 `json:"executor_number,omitempty"`

	ExecutorCuNumber *int32 `json:"executor_cu_number,omitempty"`

	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`
}

func (o FlinkJobConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobConfigInfo struct{}"
	}

	return strings.Join([]string{"FlinkJobConfigInfo", string(data)}, " ")
}

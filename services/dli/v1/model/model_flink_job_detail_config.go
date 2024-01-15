package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkJobDetailConfig struct {
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

	RuntimeConfig *string `json:"runtime_config,omitempty"`

	GraphEditorEnabled *bool `json:"graph_editor_enabled,omitempty"`

	GraphEditorData *string `json:"graph_editor_data,omitempty"`

	ResumeMaxNum *int64 `json:"resume_max_num,omitempty"`

	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	ConfigUrl *string `json:"config_url,omitempty"`

	TmCus *int32 `json:"tm_cus,omitempty"`

	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	Image *string `json:"image,omitempty"`

	Feature *string `json:"feature,omitempty"`

	FlinkVersion *string `json:"flink_version,omitempty"`

	OperatorConfig *string `json:"operator_config,omitempty"`

	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`
}

func (o FlinkJobDetailConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobDetailConfig struct{}"
	}

	return strings.Join([]string{"FlinkJobDetailConfig", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkSqlJobRequestBody struct {
	Name *string `json:"name,omitempty"`

	Desc *string `json:"desc,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	SqlBody *string `json:"sql_body,omitempty"`

	RunMode *string `json:"run_mode,omitempty"`

	CuNumber *int32 `json:"cu_number,omitempty"`

	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	CheckpointEnabled *bool `json:"checkpoint_enabled,omitempty"`

	CheckpointMode *int32 `json:"checkpoint_mode,omitempty"`

	CheckpointInterval *int32 `json:"checkpoint_interval,omitempty"`

	ObsBucket *string `json:"obs_bucket,omitempty"`

	LogEnabled *bool `json:"log_enabled,omitempty"`

	SmnTopic *string `json:"smn_topic,omitempty"`

	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	IdleStateRetention *int32 `json:"idle_state_retention,omitempty"`

	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	DirtyDataStrategy *string `json:"dirty_data_strategy,omitempty"`

	UdfJarUrl *string `json:"udf_jar_url,omitempty"`

	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	TmCus *int32 `json:"tm_cus,omitempty"`

	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	ResumeMaxNum *int32 `json:"resume_max_num,omitempty"`

	RuntimeConfig *string `json:"runtime_config,omitempty"`

	OperatorConfig *string `json:"operator_config,omitempty"`

	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`

	FlinkVersion *string `json:"flink_version,omitempty"`
}

func (o UpdateFlinkSqlJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobRequestBody", string(data)}, " ")
}

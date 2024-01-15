package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkJarJobRequestBody struct {
	Name string `json:"name"`

	Desc *string `json:"desc,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	CuNumber *int32 `json:"cu_number,omitempty"`

	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	LogEnabled *bool `json:"log_enabled,omitempty"`

	ObsBucket *string `json:"obs_bucket,omitempty"`

	SmnTopic *string `json:"smn_topic,omitempty"`

	MainClass *string `json:"main_class,omitempty"`

	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	Entrypoint *string `json:"entrypoint,omitempty"`

	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	FlinkVersion *string `json:"flink_version,omitempty"`

	Image *string `json:"image,omitempty"`

	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	TmCus *int32 `json:"tm_cus,omitempty"`

	Feature *string `json:"feature,omitempty"`

	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	ResumeMaxNum *int32 `json:"resume_max_num,omitempty"`

	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	RuntimeConfig *string `json:"runtime_config,omitempty"`

	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o CreateFlinkJarJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarJobRequestBody", string(data)}, " ")
}

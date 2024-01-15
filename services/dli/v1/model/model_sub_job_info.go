package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SubJobInfo struct {
	Id *int32 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	SubmissionTime *string `json:"submission_time,omitempty"`

	CompletionTime *string `json:"completion_time,omitempty"`

	StageIds *[]int32 `json:"stage_ids,omitempty"`

	JobGroup *string `json:"job_group,omitempty"`

	Status *string `json:"status,omitempty"`

	NumTasks *int32 `json:"num_tasks,omitempty"`

	NumActiveTasks *int32 `json:"num_active_tasks,omitempty"`

	NumCompletedTasks *int32 `json:"num_completed_tasks,omitempty"`

	NumSkippedTasks *int32 `json:"num_skipped_tasks,omitempty"`

	NumFailedTasks *int32 `json:"num_failed_tasks,omitempty"`

	NumKilledTasks *int32 `json:"num_killed_tasks,omitempty"`

	NumCompletedIndices *int32 `json:"num_completed_indices,omitempty"`

	NumActiveStages *int32 `json:"num_active_stages,omitempty"`

	NumCompletedStages *int32 `json:"num_completed_stages,omitempty"`

	NumSkippedStages *int32 `json:"num_skipped_stages,omitempty"`

	NumFailedStages *int32 `json:"num_failed_stages,omitempty"`

	KilledTasksSummary map[string]int32 `json:"killed_tasks_summary,omitempty"`
}

func (o SubJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobInfo struct{}"
	}

	return strings.Join([]string{"SubJobInfo", string(data)}, " ")
}

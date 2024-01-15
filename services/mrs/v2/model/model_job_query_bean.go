package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobQueryBean struct {
	JobId *string `json:"job_id,omitempty"`

	User *string `json:"user,omitempty"`

	JobName *string `json:"job_name,omitempty"`

	JobResult *string `json:"job_result,omitempty"`

	JobState *string `json:"job_state,omitempty"`

	JobProgress *float32 `json:"job_progress,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	StartedTime *int64 `json:"started_time,omitempty"`

	SubmittedTime *int64 `json:"submitted_time,omitempty"`

	FinishedTime *int64 `json:"finished_time,omitempty"`

	ElapsedTime *int64 `json:"elapsed_time,omitempty"`

	Arguments *string `json:"arguments,omitempty"`

	LauncherId *string `json:"launcher_id,omitempty"`

	Properties *string `json:"properties,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	TrackingUrl *string `json:"tracking_url,omitempty"`

	Queue *string `json:"queue,omitempty"`
}

func (o JobQueryBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobQueryBean struct{}"
	}

	return strings.Join([]string{"JobQueryBean", string(data)}, " ")
}

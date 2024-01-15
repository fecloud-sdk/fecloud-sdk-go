package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowJobExeListNewRequest struct {
	ClusterId string `json:"cluster_id"`

	JobName *string `json:"job_name,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	User *string `json:"user,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	JobState *ShowJobExeListNewRequestJobState `json:"job_state,omitempty"`

	JobResult *ShowJobExeListNewRequestJobResult `json:"job_result,omitempty"`

	Queue *string `json:"queue,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	SortBy *string `json:"sort_by,omitempty"`

	SubmittedTimeBegin *int64 `json:"submitted_time_begin,omitempty"`

	SubmittedTimeEnd *int64 `json:"submitted_time_end,omitempty"`
}

func (o ShowJobExeListNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobExeListNewRequest struct{}"
	}

	return strings.Join([]string{"ShowJobExeListNewRequest", string(data)}, " ")
}

type ShowJobExeListNewRequestJobState struct {
	value string
}

type ShowJobExeListNewRequestJobStateEnum struct {
	FAILED     ShowJobExeListNewRequestJobState
	KILLED     ShowJobExeListNewRequestJobState
	NEW        ShowJobExeListNewRequestJobState
	NEW_SAVING ShowJobExeListNewRequestJobState
	SUBMITTED  ShowJobExeListNewRequestJobState
	ACCEPTED   ShowJobExeListNewRequestJobState
	RUNNING    ShowJobExeListNewRequestJobState
	FINISHED   ShowJobExeListNewRequestJobState
}

func GetShowJobExeListNewRequestJobStateEnum() ShowJobExeListNewRequestJobStateEnum {
	return ShowJobExeListNewRequestJobStateEnum{
		FAILED: ShowJobExeListNewRequestJobState{
			value: "FAILED",
		},
		KILLED: ShowJobExeListNewRequestJobState{
			value: "KILLED",
		},
		NEW: ShowJobExeListNewRequestJobState{
			value: "NEW",
		},
		NEW_SAVING: ShowJobExeListNewRequestJobState{
			value: "NEW_SAVING",
		},
		SUBMITTED: ShowJobExeListNewRequestJobState{
			value: "SUBMITTED",
		},
		ACCEPTED: ShowJobExeListNewRequestJobState{
			value: "ACCEPTED",
		},
		RUNNING: ShowJobExeListNewRequestJobState{
			value: "RUNNING",
		},
		FINISHED: ShowJobExeListNewRequestJobState{
			value: "FINISHED",
		},
	}
}

func (c ShowJobExeListNewRequestJobState) Value() string {
	return c.value
}

func (c ShowJobExeListNewRequestJobState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobExeListNewRequestJobState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowJobExeListNewRequestJobResult struct {
	value string
}

type ShowJobExeListNewRequestJobResultEnum struct {
	FAILED    ShowJobExeListNewRequestJobResult
	KILLED    ShowJobExeListNewRequestJobResult
	UNDEFINED ShowJobExeListNewRequestJobResult
	SUCCEEDED ShowJobExeListNewRequestJobResult
}

func GetShowJobExeListNewRequestJobResultEnum() ShowJobExeListNewRequestJobResultEnum {
	return ShowJobExeListNewRequestJobResultEnum{
		FAILED: ShowJobExeListNewRequestJobResult{
			value: "FAILED",
		},
		KILLED: ShowJobExeListNewRequestJobResult{
			value: "KILLED",
		},
		UNDEFINED: ShowJobExeListNewRequestJobResult{
			value: "UNDEFINED",
		},
		SUCCEEDED: ShowJobExeListNewRequestJobResult{
			value: "SUCCEEDED",
		},
	}
}

func (c ShowJobExeListNewRequestJobResult) Value() string {
	return c.value
}

func (c ShowJobExeListNewRequestJobResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobExeListNewRequestJobResult) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type JobSubmitResult struct {
	JobId *string `json:"job_id,omitempty"`

	State *JobSubmitResultState `json:"state,omitempty"`
}

func (o JobSubmitResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSubmitResult struct{}"
	}

	return strings.Join([]string{"JobSubmitResult", string(data)}, " ")
}

type JobSubmitResultState struct {
	value string
}

type JobSubmitResultStateEnum struct {
	COMPLETE JobSubmitResultState
	FAILED   JobSubmitResultState
}

func GetJobSubmitResultStateEnum() JobSubmitResultStateEnum {
	return JobSubmitResultStateEnum{
		COMPLETE: JobSubmitResultState{
			value: "COMPLETE",
		},
		FAILED: JobSubmitResultState{
			value: "FAILED",
		},
	}
}

func (c JobSubmitResultState) Value() string {
	return c.value
}

func (c JobSubmitResultState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobSubmitResultState) UnmarshalJSON(b []byte) error {
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

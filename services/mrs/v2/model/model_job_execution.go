package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobExecution struct {
	JobType string `json:"job_type"`

	JobName string `json:"job_name"`

	Arguments *[]string `json:"arguments,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`
}

func (o JobExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobExecution struct{}"
	}

	return strings.Join([]string{"JobExecution", string(data)}, " ")
}

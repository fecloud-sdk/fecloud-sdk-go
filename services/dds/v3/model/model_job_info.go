package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobInfo struct {
	Id string `json:"id"`

	Name string `json:"name"`

	InstanceId string `json:"instance_id"`

	InstanceName string `json:"instance_name"`

	Status string `json:"status"`

	Progress string `json:"progress"`

	FailReason string `json:"fail_reason"`

	CreatedAt string `json:"created_at"`

	EndedAt string `json:"ended_at"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

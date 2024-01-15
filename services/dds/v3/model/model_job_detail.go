package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobDetail struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status string `json:"status"`

	Created string `json:"created"`

	Ended string `json:"ended"`

	Progress string `json:"progress"`

	Instance *JobInstanceInfo `json:"instance"`

	FailReason string `json:"fail_reason"`
}

func (o JobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}

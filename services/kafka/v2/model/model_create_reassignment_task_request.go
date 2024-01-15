package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateReassignmentTaskRequest struct {
	InstanceId string `json:"instance_id"`

	Body *PartitionReassignRequest `json:"body,omitempty"`
}

func (o CreateReassignmentTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReassignmentTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateReassignmentTaskRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExpandReplicasetNodeRequest struct {
	InstanceId string `json:"instance_id"`

	Body *EnlargeReplicasetNodeRequestBody `json:"body,omitempty"`
}

func (o ExpandReplicasetNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandReplicasetNodeRequest struct{}"
	}

	return strings.Join([]string{"ExpandReplicasetNodeRequest", string(data)}, " ")
}

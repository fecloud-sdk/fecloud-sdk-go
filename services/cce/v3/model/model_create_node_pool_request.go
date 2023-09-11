package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateNodePoolRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *NodePool `json:"body,omitempty"`
}

func (o CreateNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolRequest struct{}"
	}

	return strings.Join([]string{"CreateNodePoolRequest", string(data)}, " ")
}

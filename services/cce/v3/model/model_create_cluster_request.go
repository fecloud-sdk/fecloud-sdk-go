package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateClusterRequest struct {
	Body *Cluster `json:"body,omitempty"`
}

func (o CreateClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterRequest", string(data)}, " ")
}

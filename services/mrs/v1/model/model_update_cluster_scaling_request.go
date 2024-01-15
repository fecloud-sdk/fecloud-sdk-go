package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateClusterScalingRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *ClusterScalingReq `json:"body,omitempty"`
}

func (o UpdateClusterScalingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterScalingRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterScalingRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AwakeClusterRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o AwakeClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwakeClusterRequest struct{}"
	}

	return strings.Join([]string{"AwakeClusterRequest", string(data)}, " ")
}

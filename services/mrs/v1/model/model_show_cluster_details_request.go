package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowClusterDetailsRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailsRequest", string(data)}, " ")
}

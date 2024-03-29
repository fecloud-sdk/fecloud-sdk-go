package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowClusterRequest struct {
	ClusterId string `json:"cluster_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	Detail *string `json:"detail,omitempty"`
}

func (o ShowClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterRequest", string(data)}, " ")
}

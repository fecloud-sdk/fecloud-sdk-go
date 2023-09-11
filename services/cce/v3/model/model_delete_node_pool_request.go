package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteNodePoolRequest struct {
	ClusterId string `json:"cluster_id"`

	NodepoolId string `json:"nodepool_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`
}

func (o DeleteNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodePoolRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodePoolRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteClusterTagRequest struct {
	ClusterId string `json:"cluster_id"`

	Key string `json:"key"`
}

func (o DeleteClusterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterTagRequest", string(data)}, " ")
}

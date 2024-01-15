package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListClusterTagsRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterTagsRequest", string(data)}, " ")
}

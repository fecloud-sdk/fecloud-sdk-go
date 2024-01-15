package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateClusterTagsRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *BatchCreateClusterTagsReq `json:"body,omitempty"`
}

func (o BatchCreateClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsRequest", string(data)}, " ")
}

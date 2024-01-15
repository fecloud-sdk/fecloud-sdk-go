package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteClusterTagsRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteClusterTagsReq `json:"body,omitempty"`
}

func (o BatchDeleteClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsRequest", string(data)}, " ")
}

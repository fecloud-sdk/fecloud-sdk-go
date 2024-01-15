package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListClustersByTagsRequest struct {
	Body *ListResourceReq `json:"body,omitempty"`
}

func (o ListClustersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClustersByTagsRequest", string(data)}, " ")
}

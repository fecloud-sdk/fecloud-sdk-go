package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTransitIpsByTagsRequest struct {
	Body *ListTagResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListTransitIpsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpsByTagsRequest", string(data)}, " ")
}

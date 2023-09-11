package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListTransitIpTagsRequest Request Object
type ListTransitIpTagsRequest struct {
}

func (o ListTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpTagsRequest", string(data)}, " ")
}

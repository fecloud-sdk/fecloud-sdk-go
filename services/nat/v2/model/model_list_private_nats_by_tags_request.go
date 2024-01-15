package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateNatsByTagsRequest struct {
	Body *ListTagResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListPrivateNatsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsByTagsRequest", string(data)}, " ")
}

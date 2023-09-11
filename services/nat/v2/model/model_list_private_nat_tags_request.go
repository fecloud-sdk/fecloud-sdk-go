package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListPrivateNatTagsRequest Request Object
type ListPrivateNatTagsRequest struct {
}

func (o ListPrivateNatTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateNatTagsRequest", string(data)}, " ")
}

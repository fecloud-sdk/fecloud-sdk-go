package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListPublicipTagsRequest Request Object
type ListPublicipTagsRequest struct {
}

func (o ListPublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipTagsRequest", string(data)}, " ")
}

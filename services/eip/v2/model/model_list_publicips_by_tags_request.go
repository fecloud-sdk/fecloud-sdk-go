package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListPublicipsByTagsRequest Request Object
type ListPublicipsByTagsRequest struct {
	Body *ListPublicipsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListPublicipsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipsByTagsRequest", string(data)}, " ")
}

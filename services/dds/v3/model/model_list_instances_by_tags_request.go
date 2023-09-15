package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListInstancesByTagsRequest Request Object
type ListInstancesByTagsRequest struct {
	Body *ListInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInstancesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesByTagsRequest", string(data)}, " ")
}

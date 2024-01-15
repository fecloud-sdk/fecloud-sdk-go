package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesByResourceTagsRequest struct {
	Body *ListInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInstancesByResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesByResourceTagsRequest", string(data)}, " ")
}

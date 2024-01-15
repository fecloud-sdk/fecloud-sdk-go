package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQueryProjectResourceTagsRequest struct {
	ResourceType string `json:"resource_type"`

	ContentType *string `json:"Content-Type,omitempty"`
}

func (o ListQueryProjectResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryProjectResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListQueryProjectResourceTagsRequest", string(data)}, " ")
}

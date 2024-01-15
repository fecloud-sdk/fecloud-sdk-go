package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListResourceInstancesRequest struct {
	ResourceType string `json:"resource_type"`

	ContentType *string `json:"Content-Type,omitempty"`

	Body *QueryResourceInstanceTagsBody `json:"body,omitempty"`
}

func (o ListResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequest", string(data)}, " ")
}

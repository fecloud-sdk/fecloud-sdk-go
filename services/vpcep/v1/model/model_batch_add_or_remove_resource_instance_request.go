package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddOrRemoveResourceInstanceRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	ContentType *string `json:"Content-Type,omitempty"`

	Body *BatchAddOrRemoveResourceInstanceBody `json:"body,omitempty"`
}

func (o BatchAddOrRemoveResourceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceRequest", string(data)}, " ")
}

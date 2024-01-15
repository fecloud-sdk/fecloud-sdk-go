package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddSharedTagsRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`

	Body *BatchAddSharedTagsRequestBody `json:"body,omitempty"`
}

func (o BatchAddSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsRequest", string(data)}, " ")
}

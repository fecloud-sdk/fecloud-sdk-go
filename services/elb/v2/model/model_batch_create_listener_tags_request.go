package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateListenerTagsRequest struct {
	ListenerId string `json:"listener_id"`

	Body *BatchCreateListenerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsRequest", string(data)}, " ")
}

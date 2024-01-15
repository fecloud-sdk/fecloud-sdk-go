package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteListenerTagsRequest struct {
	ListenerId string `json:"listener_id"`

	Body *BatchDeleteListenerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsRequest", string(data)}, " ")
}

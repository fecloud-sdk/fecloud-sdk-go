package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerTagsRequest struct {
	ListenerId string `json:"listener_id"`

	Body *CreateListenerTagsRequestBody `json:"body,omitempty"`
}

func (o CreateListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsRequest", string(data)}, " ")
}

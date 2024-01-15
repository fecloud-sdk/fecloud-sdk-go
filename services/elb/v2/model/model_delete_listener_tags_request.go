package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteListenerTagsRequest struct {
	ListenerId string `json:"listener_id"`

	Key string `json:"key"`
}

func (o DeleteListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerTagsRequest", string(data)}, " ")
}

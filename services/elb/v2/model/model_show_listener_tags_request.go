package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowListenerTagsRequest Request Object
type ShowListenerTagsRequest struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`
}

func (o ShowListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerTagsRequest", string(data)}, " ")
}

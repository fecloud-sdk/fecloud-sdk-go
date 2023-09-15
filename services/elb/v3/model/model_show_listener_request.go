package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowListenerRequest Request Object
type ShowListenerRequest struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`
}

func (o ShowListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerRequest", string(data)}, " ")
}

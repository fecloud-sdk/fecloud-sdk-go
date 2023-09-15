package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteListenerRequest Request Object
type DeleteListenerRequest struct {

	// 监听器id
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerRequest", string(data)}, " ")
}

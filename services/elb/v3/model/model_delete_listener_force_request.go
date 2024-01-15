package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteListenerForceRequest struct {
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerForceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerForceRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerForceRequest", string(data)}, " ")
}

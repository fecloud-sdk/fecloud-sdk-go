package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateListenerRequest struct {
	ListenerId string `json:"listener_id"`

	Body *UpdateListenerRequestBody `json:"body,omitempty"`
}

func (o UpdateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerRequest struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequest", string(data)}, " ")
}

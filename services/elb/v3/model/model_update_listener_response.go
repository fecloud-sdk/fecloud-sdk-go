package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateListenerResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Listener       *Listener `json:"listener,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerResponse struct{}"
	}

	return strings.Join([]string{"UpdateListenerResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerResponse struct {
	Listener       *ListenerResp `json:"listener,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerResponse", string(data)}, " ")
}

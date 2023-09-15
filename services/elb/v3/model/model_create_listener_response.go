package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateListenerResponse Response Object
type CreateListenerResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Listener       *Listener `json:"listener,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerResponse", string(data)}, " ")
}

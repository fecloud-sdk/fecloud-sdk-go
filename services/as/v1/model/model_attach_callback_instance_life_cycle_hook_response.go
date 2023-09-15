package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AttachCallbackInstanceLifeCycleHookResponse Response Object
type AttachCallbackInstanceLifeCycleHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachCallbackInstanceLifeCycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachCallbackInstanceLifeCycleHookResponse struct{}"
	}

	return strings.Join([]string{"AttachCallbackInstanceLifeCycleHookResponse", string(data)}, " ")
}

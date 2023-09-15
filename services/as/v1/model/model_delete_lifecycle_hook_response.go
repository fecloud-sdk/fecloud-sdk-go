package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteLifecycleHookResponse Response Object
type DeleteLifecycleHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLifecycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLifecycleHookResponse struct{}"
	}

	return strings.Join([]string{"DeleteLifecycleHookResponse", string(data)}, " ")
}

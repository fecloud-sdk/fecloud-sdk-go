package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLifeCycleHooksResponse struct {
	LifecycleHooks *[]LifecycleHookList `json:"lifecycle_hooks,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListLifeCycleHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLifeCycleHooksResponse struct{}"
	}

	return strings.Join([]string{"ListLifeCycleHooksResponse", string(data)}, " ")
}

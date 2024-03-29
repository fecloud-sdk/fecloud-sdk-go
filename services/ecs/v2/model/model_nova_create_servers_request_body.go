package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaCreateServersRequestBody struct {
	Server *NovaCreateServersOption `json:"server"`

	OsschedulerHints *NovaCreateServersSchedulerHint `json:"os:scheduler_hints,omitempty"`
}

func (o NovaCreateServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersRequestBody struct{}"
	}

	return strings.Join([]string{"NovaCreateServersRequestBody", string(data)}, " ")
}

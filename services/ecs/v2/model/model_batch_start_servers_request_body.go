package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStartServersRequestBody struct {
	OsStart *BatchStartServersOption `json:"os-start"`
}

func (o BatchStartServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStartServersRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStartServersRequest struct {
	Body *BatchStartServersRequestBody `json:"body,omitempty"`
}

func (o BatchStartServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStartServersRequest", string(data)}, " ")
}

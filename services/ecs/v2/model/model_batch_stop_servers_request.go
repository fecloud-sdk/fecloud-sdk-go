package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStopServersRequest struct {
	Body *BatchStopServersRequestBody `json:"body,omitempty"`
}

func (o BatchStopServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStopServersRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStopServersRequestBody struct {
	OsStop *BatchStopServersOption `json:"os-stop"`
}

func (o BatchStopServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStopServersRequestBody", string(data)}, " ")
}

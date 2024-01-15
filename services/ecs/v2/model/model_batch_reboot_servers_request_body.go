package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRebootServersRequestBody struct {
	Reboot *BatchRebootSeversOption `json:"reboot"`
}

func (o BatchRebootServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRebootServersRequestBody", string(data)}, " ")
}

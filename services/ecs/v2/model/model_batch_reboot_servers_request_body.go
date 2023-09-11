package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchRebootServersRequestBody This is a auto create Body Object
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

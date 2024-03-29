package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerAutoTerminateTimeRequestBody struct {
	AutoTerminateTime string `json:"auto_terminate_time"`
}

func (o UpdateServerAutoTerminateTimeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerAutoTerminateTimeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServerAutoTerminateTimeRequestBody", string(data)}, " ")
}

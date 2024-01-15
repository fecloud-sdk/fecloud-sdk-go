package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreToExistingInstanceRequestBodyTarget struct {
	InstanceId string `json:"instance_id"`
}

func (o RestoreToExistingInstanceRequestBodyTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequestBodyTarget struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequestBodyTarget", string(data)}, " ")
}

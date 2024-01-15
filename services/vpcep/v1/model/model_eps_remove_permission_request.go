package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EpsRemovePermissionRequest struct {
	Id string `json:"id"`
}

func (o EpsRemovePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsRemovePermissionRequest struct{}"
	}

	return strings.Join([]string{"EpsRemovePermissionRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRemovePermissionRequest struct {
	Permissions []EpsRemovePermissionRequest `json:"permissions"`
}

func (o BatchRemovePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemovePermissionRequest struct{}"
	}

	return strings.Join([]string{"BatchRemovePermissionRequest", string(data)}, " ")
}

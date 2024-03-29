package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateAddonInstanceRequest struct {
	Body *InstanceRequest `json:"body,omitempty"`
}

func (o CreateAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateAddonInstanceRequest", string(data)}, " ")
}

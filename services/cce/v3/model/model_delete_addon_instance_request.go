package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAddonInstanceRequest struct {
	Id string `json:"id"`

	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o DeleteAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddonInstanceRequest", string(data)}, " ")
}

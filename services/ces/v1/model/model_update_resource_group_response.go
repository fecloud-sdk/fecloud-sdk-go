package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateResourceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupResponse", string(data)}, " ")
}

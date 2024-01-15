package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceUserResponse", string(data)}, " ")
}

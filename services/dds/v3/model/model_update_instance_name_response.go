package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameResponse", string(data)}, " ")
}

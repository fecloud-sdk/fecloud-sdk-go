package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceResponse", string(data)}, " ")
}

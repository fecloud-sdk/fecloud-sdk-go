package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceNameRequestBody struct {
	Name string `json:"name"`
}

func (o UpdateInstanceNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateNameRequestBody struct {
	NewInstanceName string `json:"new_instance_name"`
}

func (o UpdateNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNameRequestBody", string(data)}, " ")
}

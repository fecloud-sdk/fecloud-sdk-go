package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateGlobalVariableResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGlobalVariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalVariableResponse struct{}"
	}

	return strings.Join([]string{"UpdateGlobalVariableResponse", string(data)}, " ")
}

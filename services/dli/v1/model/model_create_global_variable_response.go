package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateGlobalVariableResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGlobalVariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalVariableResponse struct{}"
	}

	return strings.Join([]string{"CreateGlobalVariableResponse", string(data)}, " ")
}

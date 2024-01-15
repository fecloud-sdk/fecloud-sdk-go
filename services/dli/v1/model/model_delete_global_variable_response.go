package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteGlobalVariableResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGlobalVariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalVariableResponse struct{}"
	}

	return strings.Join([]string{"DeleteGlobalVariableResponse", string(data)}, " ")
}

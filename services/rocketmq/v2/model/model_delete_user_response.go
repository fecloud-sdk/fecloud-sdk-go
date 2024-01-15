package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteUserResponse struct {
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserResponse", string(data)}, " ")
}

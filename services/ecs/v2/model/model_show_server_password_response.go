package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowServerPasswordResponse struct {
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"ShowServerPasswordResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpResponse struct{}"
	}

	return strings.Join([]string{"CreateIpResponse", string(data)}, " ")
}

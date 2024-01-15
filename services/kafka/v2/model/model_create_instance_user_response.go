package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInstanceUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceUserResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserResponse", string(data)}, " ")
}

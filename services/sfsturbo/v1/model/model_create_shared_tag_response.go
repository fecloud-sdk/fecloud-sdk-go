package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSharedTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSharedTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSharedTagResponse", string(data)}, " ")
}

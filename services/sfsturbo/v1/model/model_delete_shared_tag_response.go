package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSharedTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSharedTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSharedTagResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEventResponse struct{}"
	}

	return strings.Join([]string{"ChangeEventResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Error struct {
	Message *string `json:"message,omitempty"`

	Code *string `json:"code,omitempty"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}

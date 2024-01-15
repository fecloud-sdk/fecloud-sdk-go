package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ActionProgress struct {
	Creating *string `json:"CREATING,omitempty"`
}

func (o ActionProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionProgress struct{}"
	}

	return strings.Join([]string{"ActionProgress", string(data)}, " ")
}

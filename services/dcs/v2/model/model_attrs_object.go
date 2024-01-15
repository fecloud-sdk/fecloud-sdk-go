package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttrsObject struct {
	Capacity *string `json:"capacity,omitempty"`

	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o AttrsObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttrsObject struct{}"
	}

	return strings.Join([]string{"AttrsObject", string(data)}, " ")
}

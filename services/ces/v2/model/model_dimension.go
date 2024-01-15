package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Dimension struct {
	Name string `json:"name"`

	Value *string `json:"value,omitempty"`
}

func (o Dimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dimension struct{}"
	}

	return strings.Join([]string{"Dimension", string(data)}, " ")
}

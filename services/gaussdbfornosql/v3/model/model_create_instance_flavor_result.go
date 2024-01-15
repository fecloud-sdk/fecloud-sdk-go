package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceFlavorResult struct {
	Num *string `json:"num,omitempty"`

	Storage *string `json:"storage,omitempty"`

	Size *string `json:"size,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`
}

func (o CreateInstanceFlavorResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceFlavorResult struct{}"
	}

	return strings.Join([]string{"CreateInstanceFlavorResult", string(data)}, " ")
}

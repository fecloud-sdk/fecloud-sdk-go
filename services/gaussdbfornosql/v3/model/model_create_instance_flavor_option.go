package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceFlavorOption struct {
	Num string `json:"num"`

	Storage string `json:"storage"`

	Size string `json:"size"`

	SpecCode string `json:"spec_code"`
}

func (o CreateInstanceFlavorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceFlavorOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceFlavorOption", string(data)}, " ")
}

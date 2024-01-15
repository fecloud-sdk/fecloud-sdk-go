package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ValuesItem struct {
	Values *[]ListApiVersionsItem `json:"values,omitempty"`
}

func (o ValuesItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValuesItem struct{}"
	}

	return strings.Join([]string{"ValuesItem", string(data)}, " ")
}

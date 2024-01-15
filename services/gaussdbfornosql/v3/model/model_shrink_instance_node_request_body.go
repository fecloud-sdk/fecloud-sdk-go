package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShrinkInstanceNodeRequestBody struct {
	Num *int32 `json:"num,omitempty"`

	NodeList *[]string `json:"node_list,omitempty"`
}

func (o ShrinkInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeRequestBody", string(data)}, " ")
}

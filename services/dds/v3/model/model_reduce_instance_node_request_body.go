package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ReduceInstanceNodeRequestBody struct {
	Num *int32 `json:"num,omitempty"`

	NodeList *[]string `json:"node_list,omitempty"`
}

func (o ReduceInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ReduceInstanceNodeRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociaterouterRequestBody struct {
	Router *Router `json:"router"`
}

func (o DisassociaterouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociaterouterRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociaterouterRequestBody", string(data)}, " ")
}

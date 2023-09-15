package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreInstanceRequestBody struct {
	Source *Source `json:"source"`

	Target *Target `json:"target"`
}

func (o RestoreInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequestBody", string(data)}, " ")
}

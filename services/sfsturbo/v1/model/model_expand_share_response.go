package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExpandShareResponse struct {
	Id *string `json:"id,omitempty"`

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandShareResponse struct{}"
	}

	return strings.Join([]string{"ExpandShareResponse", string(data)}, " ")
}

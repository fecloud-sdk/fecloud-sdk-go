package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Link struct {
	Href *string `json:"href,omitempty"`

	Rel *string `json:"rel,omitempty"`
}

func (o Link) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}

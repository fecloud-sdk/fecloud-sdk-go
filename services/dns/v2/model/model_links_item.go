package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LinksItem struct {
	Href string `json:"href"`

	Rel string `json:"rel"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}

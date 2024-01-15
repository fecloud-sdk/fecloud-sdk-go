package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LinksItem struct {
	FileName *string `json:"file_name,omitempty"`

	Link *string `json:"link,omitempty"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}

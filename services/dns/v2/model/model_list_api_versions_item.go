package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListApiVersionsItem struct {
	Status *string `json:"status,omitempty"`

	Id *string `json:"id,omitempty"`

	Links *[]LinksItem `json:"links,omitempty"`
}

func (o ListApiVersionsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsItem struct{}"
	}

	return strings.Join([]string{"ListApiVersionsItem", string(data)}, " ")
}

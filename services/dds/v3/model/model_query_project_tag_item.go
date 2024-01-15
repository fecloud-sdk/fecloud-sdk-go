package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryProjectTagItem struct {
	Key *string `json:"key,omitempty"`

	Values *[]string `json:"values,omitempty"`
}

func (o QueryProjectTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryProjectTagItem struct{}"
	}

	return strings.Join([]string{"QueryProjectTagItem", string(data)}, " ")
}

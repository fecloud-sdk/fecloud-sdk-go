package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TagResponse struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o TagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResponse struct{}"
	}

	return strings.Join([]string{"TagResponse", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourceList struct {
	Id string `json:"id"`
}

func (o ResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceList struct{}"
	}

	return strings.Join([]string{"ResourceList", string(data)}, " ")
}

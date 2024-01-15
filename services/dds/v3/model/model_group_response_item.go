package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GroupResponseItem struct {
	Type string `json:"type"`

	Id string `json:"id"`

	Name string `json:"name"`

	Status string `json:"status"`

	Volume *Volume `json:"volume"`

	Nodes []NodeItem `json:"nodes"`
}

func (o GroupResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupResponseItem struct{}"
	}

	return strings.Join([]string{"GroupResponseItem", string(data)}, " ")
}

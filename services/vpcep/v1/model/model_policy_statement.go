package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PolicyStatement struct {
	Effect string `json:"Effect"`

	Action []string `json:"Action"`

	Resource []string `json:"Resource"`
}

func (o PolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatement struct{}"
	}

	return strings.Join([]string{"PolicyStatement", string(data)}, " ")
}

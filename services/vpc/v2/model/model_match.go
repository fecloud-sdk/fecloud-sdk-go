package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Match struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ErrorlogResult struct {
	NodeName string `json:"node_name"`

	Level string `json:"level"`

	Time string `json:"time"`

	Content string `json:"content"`
}

func (o ErrorlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorlogResult struct{}"
	}

	return strings.Join([]string{"ErrorlogResult", string(data)}, " ")
}

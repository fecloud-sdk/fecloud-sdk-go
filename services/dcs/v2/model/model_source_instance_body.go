package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SourceInstanceBody struct {
	Addrs string `json:"addrs"`

	Password *string `json:"password,omitempty"`
}

func (o SourceInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInstanceBody struct{}"
	}

	return strings.Join([]string{"SourceInstanceBody", string(data)}, " ")
}

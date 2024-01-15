package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Hash struct {
}

func (o Hash) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hash struct{}"
	}

	return strings.Join([]string{"Hash", string(data)}, " ")
}

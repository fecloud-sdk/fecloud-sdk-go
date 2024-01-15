package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UnprotectHostNum struct {
}

func (o UnprotectHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnprotectHostNum struct{}"
	}

	return strings.Join([]string{"UnprotectHostNum", string(data)}, " ")
}

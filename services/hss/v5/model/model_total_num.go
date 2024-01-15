package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TotalNum struct {
}

func (o TotalNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalNum struct{}"
	}

	return strings.Join([]string{"TotalNum", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OperateType struct {
}

func (o OperateType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateType struct{}"
	}

	return strings.Join([]string{"OperateType", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OsType struct {
}

func (o OsType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsType struct{}"
	}

	return strings.Join([]string{"OsType", string(data)}, " ")
}

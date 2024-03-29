package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MysqlReadOnlySwitch struct {
	Readonly bool `json:"readonly"`
}

func (o MysqlReadOnlySwitch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlReadOnlySwitch struct{}"
	}

	return strings.Join([]string{"MysqlReadOnlySwitch", string(data)}, " ")
}

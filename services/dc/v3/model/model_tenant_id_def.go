package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TenantIdDef struct {
}

func (o TenantIdDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantIdDef struct{}"
	}

	return strings.Join([]string{"TenantIdDef", string(data)}, " ")
}

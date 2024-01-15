package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnterpriseProjectName struct {
}

func (o EnterpriseProjectName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectName struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectName", string(data)}, " ")
}

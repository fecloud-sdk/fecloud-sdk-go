package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UserWithPrivilege struct {
	Name string `json:"name"`

	Readonly bool `json:"readonly"`
}

func (o UserWithPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserWithPrivilege struct{}"
	}

	return strings.Join([]string{"UserWithPrivilege", string(data)}, " ")
}

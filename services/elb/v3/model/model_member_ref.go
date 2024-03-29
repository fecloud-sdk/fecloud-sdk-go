package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MemberRef struct {
	Id string `json:"id"`
}

func (o MemberRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberRef struct{}"
	}

	return strings.Join([]string{"MemberRef", string(data)}, " ")
}

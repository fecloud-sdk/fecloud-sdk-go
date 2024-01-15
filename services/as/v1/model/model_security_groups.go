package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SecurityGroups struct {
	Id string `json:"id"`
}

func (o SecurityGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroups struct{}"
	}

	return strings.Join([]string{"SecurityGroups", string(data)}, " ")
}

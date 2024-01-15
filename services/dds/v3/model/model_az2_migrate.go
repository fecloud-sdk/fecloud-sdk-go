package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Az2Migrate struct {
	Code string `json:"code"`

	Description string `json:"description"`

	Status string `json:"status"`
}

func (o Az2Migrate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Az2Migrate struct{}"
	}

	return strings.Join([]string{"Az2Migrate", string(data)}, " ")
}

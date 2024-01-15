package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShareName struct {
	Name string `json:"name"`
}

func (o ShareName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareName struct{}"
	}

	return strings.Join([]string{"ShareName", string(data)}, " ")
}
